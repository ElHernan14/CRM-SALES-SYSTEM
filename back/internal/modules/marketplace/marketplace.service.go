package marketplace

import (
	"context"
	coreDto "crm-system-sales/internal/core/dto"
	errorHandler "crm-system-sales/internal/core/error"
	tenantHelper "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/modules/company"
	invoiceconstants "crm-system-sales/internal/modules/invoice/constants"
	invoicemodel "crm-system-sales/internal/modules/invoice/models"
	invoicerepo "crm-system-sales/internal/modules/invoice/repository"
	marketplacedto "crm-system-sales/internal/modules/marketplace/dto"
	"net/http"
)

type MarketplaceService interface {
	GetSuppliers(ctx context.Context, req *marketplacedto.GetSuppliersRequest) (*marketplacedto.GetSuppliersResponse, error)
	GetSupplierByID(ctx context.Context, supplierID int) (*marketplacedto.SupplierResponse, error)
	EnsureCart(ctx context.Context, req *marketplacedto.EnsureMarketplaceCartRequest) (*marketplacedto.EnsureMarketplaceCartResponse, error)
}

type marketplaceService struct {
	repo        MarketplaceRepository
	invoiceRepo invoicerepo.InvoiceRepository
	companyRepo company.CompanyRepository
}

func NewMarketplaceService(repo MarketplaceRepository, invoiceRepo invoicerepo.InvoiceRepository, companyRepo company.CompanyRepository) MarketplaceService {
	return &marketplaceService{repo: repo, invoiceRepo: invoiceRepo, companyRepo: companyRepo}
}

func (s *marketplaceService) GetSuppliers(ctx context.Context, req *marketplacedto.GetSuppliersRequest) (*marketplacedto.GetSuppliersResponse, error) {
	tenant := tenantHelper.GetTenant(ctx)
	var excludedCompanyID *int
	if tenant != nil {
		excludedCompanyID = tenant.CompanyID
	}

	items, total, err := s.repo.ListSuppliers(ctx, req, excludedCompanyID)
	if err != nil {
		return nil, err
	}

	return &marketplacedto.GetSuppliersResponse{
		Items: items,
		Meta:  coreDto.NewMeta(req.Page, req.Limit, total),
	}, nil
}

func (s *marketplaceService) GetSupplierByID(ctx context.Context, supplierID int) (*marketplacedto.SupplierResponse, error) {
	if supplierID <= 0 {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "supplier_id invalido")
	}

	tenant := tenantHelper.GetTenant(ctx)
	var excludedCompanyID *int
	if tenant != nil {
		excludedCompanyID = tenant.CompanyID
	}

	supplier, err := s.repo.GetSupplierByID(ctx, supplierID, excludedCompanyID)
	if err != nil {
		return nil, err
	}
	if supplier == nil {
		return nil, errorHandler.NewAppError(http.StatusNotFound, "supplier no encontrado")
	}

	return supplier, nil
}

func (s *marketplaceService) EnsureCart(ctx context.Context, req *marketplacedto.EnsureMarketplaceCartRequest) (*marketplacedto.EnsureMarketplaceCartResponse, error) {
	tenant := tenantHelper.GetTenant(ctx)
	if tenant == nil || tenant.CompanyID == nil || tenant.ClientID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "solo empresas pueden crear compras B2B")
	}

	if req.SellerCompanyID <= 0 {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "seller_company_id invalido")
	}

	if req.SellerCompanyID == *tenant.CompanyID {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "una empresa no puede comprarse a si misma")
	}

	seller, err := s.companyRepo.GetByID(ctx, req.SellerCompanyID)
	if err != nil {
		return nil, err
	}
	if seller == nil {
		return nil, errorHandler.NewAppError(http.StatusNotFound, "empresa proveedora no encontrada")
	}

	draft, err := s.invoiceRepo.GetActiveDraft(ctx, *tenant.ClientID, req.SellerCompanyID, invoiceconstants.InvoiceSourceERP)
	if err != nil {
		return nil, err
	}

	if draft == nil {
		draft = &invoicemodel.Invoice{
			BuyerClientID:   *tenant.ClientID,
			SellerCompanyID: req.SellerCompanyID,
			CreatedByUserID: tenant.UserID,
			StatusInvoice:   invoiceconstants.InvoiceDraft,
			Source:          invoiceconstants.InvoiceSourceERP,
			TotalAmount:     0,
		}

		if err := s.invoiceRepo.Create(ctx, nil, draft); err != nil {
			return nil, errorHandler.NewAppError(http.StatusInternalServerError, "error creando carrito B2B")
		}
	}

	return &marketplacedto.EnsureMarketplaceCartResponse{
		InvoiceID:       draft.ID,
		BuyerClientID:   draft.BuyerClientID,
		SellerCompanyID: draft.SellerCompanyID,
		StatusInvoice:   draft.StatusInvoice,
		Source:          draft.Source,
	}, nil
}
