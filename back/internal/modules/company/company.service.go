package company

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	constants "crm-system-sales/internal/constants"
	access "crm-system-sales/internal/core/access"
	metadto "crm-system-sales/internal/core/dto"
	errorHandler "crm-system-sales/internal/core/error"
	tenant "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/core/transaction"
	company "crm-system-sales/internal/models/company"
	userModel "crm-system-sales/internal/models/users"
	"crm-system-sales/internal/modules/auth"
	companydto "crm-system-sales/internal/modules/company/dto"
	"crm-system-sales/internal/modules/users"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type CompanyService interface {
	CreateCompany(ctx context.Context, req *companydto.CreateCompanyRequest) (*companydto.CompanyResponse, error)
	GetMyCompany(ctx context.Context) (*companydto.CompanyResponse, error)
	GetCompanies(ctx context.Context, req *companydto.GetCompaniesRequest) (*companydto.GetCompaniesResponse, error)
	GetByID(ctx context.Context, id int) (*companydto.CompanyResponse, error)
	UploadLogo(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*companydto.UploadCompanyLogoResponse, error)
}

type companyService struct {
	db        *sql.DB
	repo      CompanyRepository
	authRepo  auth.AuthRepository
	userRepo  users.UserRepository
	uploadDir string
}

func NewCompanyService(db *sql.DB, repo CompanyRepository, authRepo auth.AuthRepository, userRepo users.UserRepository, uploadDir string) CompanyService {
	return &companyService{db: db, repo: repo, authRepo: authRepo, userRepo: userRepo, uploadDir: uploadDir}
}

func (s *companyService) CreateCompany(ctx context.Context, req *companydto.CreateCompanyRequest) (*companydto.CompanyResponse, error) {
	var err error

	exists, err := s.userRepo.EmailExists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorHandler.ErrConflict
	}

	var resp *companydto.CompanyResponse
	err = transaction.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {
		company := &company.Company{Name: req.Name, Description: req.Description, Status: 1}

		company, err = s.repo.Create(tx, company)
		if err != nil {
			return err
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user := &userModel.User{Email: req.Email, PasswordHash: string(hash), CompanyID: &company.ID, Status: 1}

		user, err = s.userRepo.Create(tx, user)
		if err != nil {
			return err
		}

		if err := s.authRepo.AssignRole(tx, user.ID, constants.RoleManager); err != nil {
			return err
		}

		resp = &companydto.CompanyResponse{ID: company.ID, Name: company.Name, Description: company.Description}
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

func (s *companyService) GetCompanies(ctx context.Context, req *companydto.GetCompaniesRequest) (*companydto.GetCompaniesResponse, error) {
	offset := (req.Page - 1) * req.Limit
	companies, total, err := s.repo.GetAll(ctx, req.Search, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	data := make([]companydto.CompanyListItem, 0, len(companies))
	for _, c := range companies {
		data = append(data, companydto.CompanyListItem{ID: c.ID, Name: c.Name, Logo: c.LogoPath})
	}

	return &companydto.GetCompaniesResponse{Data: data, Meta: metadto.Meta{Page: req.Page, Limit: req.Limit, Total: total}}, nil
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

	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	if !allowed[ext] {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "formato de imagen inválido")
	}

	companyDir := filepath.Join(s.uploadDir, "company", "logos")
	if err := os.MkdirAll(companyDir, 0755); err != nil {
		return nil, err
	}

	filename := fmt.Sprintf("company_%d_%d%s", *tenant.CompanyID, time.Now().UnixNano(), ext)
	absolutePath := filepath.Join(companyDir, filename)

	dst, err := os.Create(absolutePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return nil, err
	}

	publicPath := "/uploads/company/logos/" + filename
	if err := s.repo.UpdateLogo(ctx, *tenant.CompanyID, publicPath); err != nil {
		if err == sql.ErrNoRows {
			return nil, errorHandler.NewAppError(http.StatusNotFound, "empresa no encontrada")
		}
		return nil, err
	}

	return &companydto.UploadCompanyLogoResponse{Logo: publicPath}, nil
}

func mapCompanyResponse(c *company.Company) *companydto.CompanyResponse {
	return &companydto.CompanyResponse{ID: c.ID, Name: c.Name, Description: c.Description, Logo: c.LogoPath, Status: &c.Status}
}
