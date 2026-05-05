package company

import (
	"context"
	constants "crm-system-sales/internal/constants"
	access "crm-system-sales/internal/core/access"
	metadto "crm-system-sales/internal/core/dto"
	errorHandler "crm-system-sales/internal/core/error"
	tenant "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/core/transaction"
	"crm-system-sales/internal/core/utils"
	company "crm-system-sales/internal/models/company"
	userModel "crm-system-sales/internal/models/users"
	"crm-system-sales/internal/modules/auth"
	companydto "crm-system-sales/internal/modules/company/dto"
	"crm-system-sales/internal/modules/users"
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type CompanyService interface {
	CreateCompany(ctx context.Context, req *companydto.CreateCompanyRequest) (*companydto.CompanyResponse, error)
	GetMyCompany(ctx context.Context) (*companydto.CompanyResponse, error)
	GetCompanies(ctx context.Context, req *companydto.GetCompaniesRequest) (*companydto.GetCompaniesResponse, error)
	GetByID(ctx context.Context, id int) (*companydto.CompanyResponse, error)
}

type companyService struct {
	db       *sql.DB
	repo     CompanyRepository
	authRepo auth.AuthRepository
	userRepo users.UserRepository
}

func NewCompanyService(db *sql.DB, repo CompanyRepository, authRepo auth.AuthRepository, userRepo users.UserRepository) CompanyService {
	return &companyService{
		db:       db,
		repo:     repo,
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *companyService) CreateCompany(
	ctx context.Context,
	req *companydto.CreateCompanyRequest,
) (*companydto.CompanyResponse, error) {

	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE CreateCompany")(err)
	}()

	exists, err := s.userRepo.EmailExists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorHandler.ErrConflict
	}

	var resp *companydto.CompanyResponse

	err = transaction.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {

		company := &company.Company{
			Name:   req.Name,
			Status: 1,
		}

		company, err = s.repo.Create(tx, company)
		if err != nil {
			return err
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		user := &userModel.User{
			Email:        req.Email,
			PasswordHash: string(hash),
			CompanyID:    &company.ID,
			Status:       1,
		}

		user, err = s.userRepo.Create(tx, user)
		if err != nil {
			return err
		}

		if err := s.authRepo.AssignRole(tx, user.ID, constants.RoleManager); err != nil {
			return err
		}

		resp = &companydto.CompanyResponse{
			ID:   company.ID,
			Name: company.Name,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *companyService) GetMyCompany(ctx context.Context) (*companydto.CompanyResponse, error) {

	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE GetMyCompany")(err)
	}()

	tenant := tenant.GetTenant(ctx)

	if tenant == nil || tenant.CompanyID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "usuario no pertenece a una empresa")
	}

	company, err := s.repo.GetByID(ctx, *tenant.CompanyID)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, errorHandler.NewAppError(http.StatusNotFound, "empresa no encontrada")
	}

	return &companydto.CompanyResponse{
		ID:   company.ID,
		Name: company.Name,
	}, nil
}

func (s *companyService) GetCompanies(
	ctx context.Context,
	req *companydto.GetCompaniesRequest,
) (*companydto.GetCompaniesResponse, error) {

	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE GetCompanies")(err)
	}()

	offset := (req.Page - 1) * req.Limit

	companies, total, err := s.repo.GetAll(ctx, req.Search, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	var data []companydto.CompanyListItem

	for _, c := range companies {
		data = append(data, companydto.CompanyListItem{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return &companydto.GetCompaniesResponse{
		Data: data,
		Meta: metadto.Meta{
			Page:  req.Page,
			Limit: req.Limit,
			Total: total,
		},
	}, nil
}

func (s *companyService) GetByID(ctx context.Context, id int) (*companydto.CompanyResponse, error) {

	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE GetCompanyByID")(err)
	}()

	tenant := tenant.GetTenant(ctx)

	if tenant == nil {
		return nil, errorHandler.ErrForbidden
	}

	companyID, err := access.ResolveCompanyScope(tenant, &id)
	if err != nil {
		return nil, err
	}

	company, err := s.repo.GetByID(ctx, *companyID)
	if err != nil {
		return nil, err
	}
	if company == nil {
		err = errorHandler.NewAppError(http.StatusNotFound, "empresa no encontrada")
		return nil, err
	}

	return &companydto.CompanyResponse{
		ID:   company.ID,
		Name: company.Name,
	}, nil
}
