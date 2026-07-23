package auth

import (
	"context"
	"crm-system-sales/internal/constants"
	authcore "crm-system-sales/internal/core/auth"
	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/transaction"
	models "crm-system-sales/internal/models/auth"
	categoryCompanyModel "crm-system-sales/internal/models/category_company"
	clientModel "crm-system-sales/internal/models/client"
	companyModel "crm-system-sales/internal/models/company"
	userModel "crm-system-sales/internal/models/users"
	authdto "crm-system-sales/internal/modules/auth/dto"
	"crm-system-sales/internal/modules/users"
	"database/sql"
	"log"
	"net/http"
)

type AuthService interface {
	Login(r *http.Request, email, password string) (*models.UserLogin, error)
	RegisterPersonal(ctx context.Context, req *authdto.PersonalRegisterRequest) (*authdto.RegisterResponse, error)
	RegisterBusiness(ctx context.Context, req *authdto.BusinessRegisterRequest) (*authdto.RegisterResponse, error)
}

type authService struct {
	db           *sql.DB
	UserRepo     users.UserRepository
	clientRepo   registerClientRepository
	companyRepo  registerCompanyRepository
	authRepo     AuthRepository
	categoryRepo registerCategoryCompanyRepository
}

type registerClientRepository interface {
	Create(tx *sql.Tx, client *clientModel.Client) error
	EmailExists(email string) (bool, error)
}

type registerCompanyRepository interface {
	Create(tx *sql.Tx, c *companyModel.Company) (*companyModel.Company, error)
}

type registerCategoryCompanyRepository interface {
	GetByID(ctx context.Context, id int) (*categoryCompanyModel.CategoryCompany, error)
}

func NewAuthService(db *sql.DB, userRepo users.UserRepository, clientRepo registerClientRepository, companyRepo registerCompanyRepository, authRepo AuthRepository, categoryRepo registerCategoryCompanyRepository) AuthService {
	return &authService{db: db, UserRepo: userRepo, clientRepo: clientRepo, companyRepo: companyRepo, authRepo: authRepo, categoryRepo: categoryRepo}
}

func (s *authService) Login(r *http.Request, email, password string) (*models.UserLogin, error) {
	var err error

	user, err := s.UserRepo.GetUserLogin(email)
	if err != nil {
		log.Println("Error fetching user for login: ", err)
		err = errorHandler.NewAppError(http.StatusUnauthorized, "invalid credentials")
		return nil, err
	}

	err = authcore.CheckPassword(password, user.PasswordHash)
	if err != nil {
		log.Println("Error checking password: ", err)
		err = errorHandler.NewAppError(http.StatusUnauthorized, "invalid credentials")
		return nil, err
	}

	return user, nil
}

func (s *authService) RegisterPersonal(ctx context.Context, req *authdto.PersonalRegisterRequest) (*authdto.RegisterResponse, error) {
	if err := s.ensureEmailAvailable(req.Email); err != nil {
		return nil, err
	}

	err := transaction.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {
		user, err := s.createUser(tx, req.Email, req.Password)
		if err != nil {
			return err
		}

		client := &clientModel.Client{
			UserID:    &user.ID,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
			Status:    1,
		}
		if err := s.clientRepo.Create(tx, client); err != nil {
			return err
		}

		return s.authRepo.AssignRole(tx, user.ID, constants.RoleUser)
	})
	if err != nil {
		return nil, err
	}

	return s.buildRegisterResponse(req.Email)
}

func (s *authService) RegisterBusiness(ctx context.Context, req *authdto.BusinessRegisterRequest) (*authdto.RegisterResponse, error) {
	if err := s.ensureEmailAvailable(req.Email); err != nil {
		return nil, err
	}

	category, err := s.categoryRepo.GetByID(ctx, req.Company.CategoryID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "Categoria de empresa no encontrada")
	}

	err = transaction.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {
		company := &companyModel.Company{
			Name:        req.Company.Name,
			CategoryID:  req.Company.CategoryID,
			Category:    category.Name,
			Description: req.Company.Description,
			Status:      1,
		}
		company, err = s.companyRepo.Create(tx, company)
		if err != nil {
			return err
		}

		user, err := s.createUser(tx, req.Email, req.Password)
		if err != nil {
			return err
		}

		client := &clientModel.Client{
			UserID:    &user.ID,
			CompanyID: &company.ID,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
			Status:    1,
		}
		if err := s.clientRepo.Create(tx, client); err != nil {
			return err
		}

		return s.authRepo.AssignRole(tx, user.ID, constants.RoleManager)
	})
	if err != nil {
		return nil, err
	}

	return s.buildRegisterResponse(req.Email)
}

func (s *authService) ensureEmailAvailable(email string) error {
	exists, err := s.UserRepo.EmailExists(email)
	if err != nil {
		return err
	}
	if exists {
		return errorHandler.NewAppError(http.StatusConflict, "email already exists")
	}

	exists, err = s.clientRepo.EmailExists(email)
	if err != nil {
		return err
	}
	if exists {
		return errorHandler.NewAppError(http.StatusConflict, "email already exists")
	}

	return nil
}

func (s *authService) createUser(tx *sql.Tx, email string, password string) (*userModel.User, error) {
	hash, err := authcore.HashPassword(password)
	if err != nil {
		return nil, errorHandler.NewAppError(http.StatusInternalServerError, "Error generando password hash")
	}

	user := &userModel.User{
		Email:        email,
		PasswordHash: hash,
		Status:       1,
	}

	return s.UserRepo.Create(tx, user)
}

func (s *authService) buildRegisterResponse(email string) (*authdto.RegisterResponse, error) {
	user, err := s.UserRepo.GetUserLogin(email)
	if err != nil {
		return nil, err
	}

	token, err := authcore.GenerateJWT(user, user.Permissions, user.Roles)
	if err != nil {
		return nil, err
	}

	return &authdto.RegisterResponse{
		Token: token,
		User: authdto.AuthUserContextResponse{
			UserID:      user.ID,
			Email:       user.Email,
			CompanyID:   user.CompanyID,
			ClientID:    user.ClientID,
			Roles:       user.Roles,
			Permissions: user.Permissions,
		},
	}, nil
}
