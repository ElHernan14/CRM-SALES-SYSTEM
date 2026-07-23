package company

import (
	"context"
	"database/sql"
	"mime/multipart"
	"net/http"

	constants "crm-system-sales/internal/constants"
	access "crm-system-sales/internal/core/access"
	metadto "crm-system-sales/internal/core/dto"
	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/files"
	tenant "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/core/transaction"
	clientModel "crm-system-sales/internal/models/client"
	company "crm-system-sales/internal/models/company"
	userModel "crm-system-sales/internal/models/users"
	"crm-system-sales/internal/modules/auth"
	categorycompany "crm-system-sales/internal/modules/category_company"
	"crm-system-sales/internal/modules/client"
	companydto "crm-system-sales/internal/modules/company/dto"
	"crm-system-sales/internal/modules/users"

	"golang.org/x/crypto/bcrypt"
)

type CompanyService interface {
	CreateCompany(ctx context.Context, req *companydto.CreateCompanyRequest) (*companydto.CompanyResponse, error)
	GetMyCompany(ctx context.Context) (*companydto.CompanyResponse, error)
	UpdateMyCompany(ctx context.Context, req *companydto.UpdateMyCompanyRequest) (*companydto.CompanyResponse, error)
	GetCompanies(ctx context.Context, req *companydto.GetCompaniesRequest) (*companydto.GetCompaniesResponse, error)
	GetByID(ctx context.Context, id int) (*companydto.CompanyResponse, error)
	UploadLogo(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*companydto.UploadCompanyLogoResponse, error)
	UploadCoverImage(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*companydto.UploadCompanyCoverImageResponse, error)
}

type companyService struct {
	db           *sql.DB
	repo         CompanyRepository
	clientRepo   client.ClientRepository
	authRepo     auth.AuthRepository
	userRepo     users.UserRepository
	categoryRepo categorycompany.CategoryCompanyRepository
	imageStorage files.ImageStorage
}

func NewCompanyService(
	db *sql.DB,
	repo CompanyRepository,
	clientRepo client.ClientRepository,
	authRepo auth.AuthRepository,
	userRepo users.UserRepository,
	categoryRepo categorycompany.CategoryCompanyRepository,
	imageStorage files.ImageStorage,
) CompanyService {
	return &companyService{
		db:           db,
		repo:         repo,
		clientRepo:   clientRepo,
		authRepo:     authRepo,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
		imageStorage: imageStorage,
	}
}

func (s *companyService) CreateCompany(ctx context.Context, req *companydto.CreateCompanyRequest) (*companydto.CompanyResponse, error) {
	exists, err := s.userRepo.EmailExists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorHandler.ErrConflict
	}

	category, err := s.categoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "Categoria de empresa no encontrada")
	}

	var resp *companydto.CompanyResponse
	err = transaction.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {
		company := &company.Company{Name: req.Name, CategoryID: req.CategoryID, Category: category.Name, Description: req.Description, Status: 1}

		company, err = s.repo.Create(tx, company)
		if err != nil {
			return err
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user := &userModel.User{Email: req.Email, PasswordHash: string(hash), Status: 1}

		user, err = s.userRepo.Create(tx, user)
		if err != nil {
			return err
		}

		companyClient := &clientModel.Client{
			UserID:    &user.ID,
			CompanyID: &company.ID,
			FirstName: company.Name,
			LastName:  "Pioneer",
			Email:     req.Email,
			Status:    1,
		}

		if err := s.clientRepo.Create(tx, companyClient); err != nil {
			return err
		}

		if err := s.authRepo.AssignRole(tx, user.ID, constants.RoleManager); err != nil {
			return err
		}

		resp = &companydto.CompanyResponse{ID: company.ID, Name: company.Name, CategoryID: company.CategoryID, Category: company.Category, Description: company.Description}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *companyService) GetMyCompany(ctx context.Context) (*companydto.CompanyResponse, error) {
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

	return mapCompanyResponse(company), nil
}

func (s *companyService) UpdateMyCompany(ctx context.Context, req *companydto.UpdateMyCompanyRequest) (*companydto.CompanyResponse, error) {
	tenant := tenant.GetTenant(ctx)
	if tenant == nil || tenant.CompanyID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "usuario no pertenece a una empresa")
	}

	var categoryName *string
	if req.CategoryID != nil {
		category, err := s.categoryRepo.GetByID(ctx, *req.CategoryID)
		if err != nil {
			return nil, err
		}
		if category == nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "categoria de empresa no encontrada")
		}
		categoryName = &category.Name
	}

	company, err := s.repo.UpdateMyCompany(ctx, *tenant.CompanyID, req.Name, req.CategoryID, categoryName, req.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorHandler.NewAppError(http.StatusNotFound, "empresa no encontrada")
		}
		return nil, err
	}

	return mapCompanyResponse(company), nil
}

func (s *companyService) GetCompanies(ctx context.Context, req *companydto.GetCompaniesRequest) (*companydto.GetCompaniesResponse, error) {
	offset := (req.Page - 1) * req.Limit
	companies, total, err := s.repo.GetAll(ctx, req.Search, req.CategoryID, req.Category, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	data := make([]companydto.CompanyListItem, 0, len(companies))
	for _, c := range companies {
		data = append(data, companydto.CompanyListItem{ID: c.ID, Name: c.Name, CategoryID: c.CategoryID, Category: c.Category, Logo: c.LogoPath, CoverImage: c.CoverImagePath})
	}

	return &companydto.GetCompaniesResponse{Data: data, Meta: metadto.NewMeta(req.Page, req.Limit, total)}, nil
}

func (s *companyService) GetByID(ctx context.Context, id int) (*companydto.CompanyResponse, error) {
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
		return nil, errorHandler.NewAppError(http.StatusNotFound, "empresa no encontrada")
	}

	return mapCompanyResponse(company), nil
}

func (s *companyService) UploadLogo(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*companydto.UploadCompanyLogoResponse, error) {
	tenant := tenant.GetTenant(ctx)
	if tenant == nil || tenant.CompanyID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "usuario no pertenece a una empresa")
	}

	publicPath, err := s.imageStorage.SaveImage(file, header, "company/logos", "company", *tenant.CompanyID)
	if err != nil {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, err.Error())
	}

	if err := s.repo.UpdateLogo(ctx, *tenant.CompanyID, publicPath); err != nil {
		if err == sql.ErrNoRows {
			return nil, errorHandler.NewAppError(http.StatusNotFound, "empresa no encontrada")
		}
		return nil, err
	}

	return &companydto.UploadCompanyLogoResponse{Logo: publicPath}, nil
}

func (s *companyService) UploadCoverImage(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*companydto.UploadCompanyCoverImageResponse, error) {
	tenant := tenant.GetTenant(ctx)
	if tenant == nil || tenant.CompanyID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "usuario no pertenece a una empresa")
	}

	publicPath, err := s.imageStorage.SaveImage(file, header, "company/covers", "company_cover", *tenant.CompanyID)
	if err != nil {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, err.Error())
	}

	if err := s.repo.UpdateCoverImage(ctx, *tenant.CompanyID, publicPath); err != nil {
		if err == sql.ErrNoRows {
			return nil, errorHandler.NewAppError(http.StatusNotFound, "empresa no encontrada")
		}
		return nil, err
	}

	return &companydto.UploadCompanyCoverImageResponse{CoverImage: publicPath}, nil
}

func mapCompanyResponse(c *company.Company) *companydto.CompanyResponse {
	return &companydto.CompanyResponse{ID: c.ID, Name: c.Name, CategoryID: c.CategoryID, Category: c.Category, Description: c.Description, Logo: c.LogoPath, CoverImage: c.CoverImagePath, Status: &c.Status}
}
