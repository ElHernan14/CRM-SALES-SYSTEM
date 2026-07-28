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
	invoiceitemrepo "crm-system-sales/internal/modules/invoice_item"
	marketplacedto "crm-system-sales/internal/modules/marketplace/dto"
	submitInvoiceWorkflow "crm-system-sales/internal/services/invoice_workflow/service"
	"net/http"
)

type MarketplaceService interface {
	GetSuppliers(ctx context.Context, req *marketplacedto.GetSuppliersRequest) (*marketplacedto.GetSuppliersResponse, error)
	GetSupplierByID(ctx context.Context, supplierID int) (*marketplacedto.SupplierResponse, error)
	EnsureCart(ctx context.Context, req *marketplacedto.EnsureMarketplaceCartRequest) (*marketplacedto.EnsureMarketplaceCartResponse, error)
	GetCart(ctx context.Context, sellerCompanyID int) (*marketplacedto.MarketplaceCartResponse, error)
	Checkout(ctx context.Context, req *marketplacedto.MarketplaceCheckoutRequest) (*marketplacedto.MarketplaceCheckoutResponse, error)
}

type marketplaceService struct {
	repo            MarketplaceRepository
	invoiceRepo     invoicerepo.InvoiceRepository
	invoiceItemRepo invoiceitemrepo.InvoiceItemRepository
	companyRepo     company.CompanyRepository
	submitWorkflow  submitInvoiceWorkflow.SubmitInvoiceWorkflow
}

func NewMarketplaceService(
	repo MarketplaceRepository,
	invoiceRepo invoicerepo.InvoiceRepository,
	invoiceItemRepo invoiceitemrepo.InvoiceItemRepository,
	companyRepo company.CompanyRepository,
	submitWorkflow submitInvoiceWorkflow.SubmitInvoiceWorkflow,
) MarketplaceService {
	return &marketplaceService{repo: repo, invoiceRepo: invoiceRepo, invoiceItemRepo: invoiceItemRepo, companyRepo: companyRepo, submitWorkflow: submitWorkflow}
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

func (s *marketplaceService) GetCart(ctx context.Context, sellerCompanyID int) (*marketplacedto.MarketplaceCartResponse, error) {
	tenant := tenantHelper.GetTenant(ctx)
	if tenant == nil || tenant.CompanyID == nil || tenant.ClientID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "solo empresas pueden consultar compras B2B")
	}

	if sellerCompanyID <= 0 {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "seller_company_id invalido")
	}

	if sellerCompanyID == *tenant.CompanyID {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "una empresa no puede comprarse a si misma")
	}

	draft, err := s.invoiceRepo.GetActiveDraft(ctx, *tenant.ClientID, sellerCompanyID, invoiceconstants.InvoiceSourceERP)
	if err != nil {
		return nil, err
	}
	if draft == nil {
		return nil, nil
	}

	cart, err := s.buildCartResponse(ctx, draft)
	if err != nil {
		return nil, err
	}
	if len(cart.Items) == 0 {
		return nil, nil
	}

	return cart, nil
}

func (s *marketplaceService) Checkout(ctx context.Context, req *marketplacedto.MarketplaceCheckoutRequest) (*marketplacedto.MarketplaceCheckoutResponse, error) {
	tenant := tenantHelper.GetTenant(ctx)
	if tenant == nil || tenant.CompanyID == nil || tenant.ClientID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "solo empresas pueden procesar compras B2B")
	}

	var invoiceID int
	if req != nil && req.InvoiceID != nil {
		invoice, err := s.invoiceRepo.GetByID(ctx, *req.InvoiceID)
		if err != nil {
			return nil, err
		}
		if invoice == nil || invoice.Source != invoiceconstants.InvoiceSourceERP {
			return nil, errorHandler.NewAppError(http.StatusNotFound, "carrito B2B no encontrado")
		}
		invoiceID = invoice.ID
	} else if req != nil && req.SellerCompanyID != nil {
		draft, err := s.invoiceRepo.GetActiveDraft(ctx, *tenant.ClientID, *req.SellerCompanyID, invoiceconstants.InvoiceSourceERP)
		if err != nil {
			return nil, err
		}
		if draft == nil {
			return nil, errorHandler.NewAppError(http.StatusNotFound, "carrito B2B no encontrado")
		}
		invoiceID = draft.ID
	} else {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "invoice_id o seller_company_id es requerido")
	}

	itemCount, err := s.invoiceItemRepo.CountByInvoice(ctx, invoiceID)
	if err != nil {
		return nil, errorHandler.NewAppError(http.StatusInternalServerError, "error validando items del carrito B2B")
	}
	if itemCount == 0 {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "el carrito B2B no posee items")
	}

	if err := s.submitWorkflow.Submit(ctx, invoiceID); err != nil {
		return nil, err
	}

	return &marketplacedto.MarketplaceCheckoutResponse{InvoiceID: invoiceID, Status: invoiceconstants.InvoicePending, Source: invoiceconstants.InvoiceSourceERP}, nil
}

func (s *marketplaceService) buildCartResponse(ctx context.Context, draft *invoicemodel.Invoice) (*marketplacedto.MarketplaceCartResponse, error) {
	seller, err := s.companyRepo.GetByID(ctx, draft.SellerCompanyID)
	if err != nil {
		return nil, err
	}
	if seller == nil {
		return nil, errorHandler.NewAppError(http.StatusNotFound, "empresa proveedora no encontrada")
	}

	items, _, err := s.invoiceItemRepo.GetByInvoiceID(ctx, draft.ID, 1, 100)
	if err != nil {
		return nil, err
	}

	res := &marketplacedto.MarketplaceCartResponse{
		InvoiceID:       draft.ID,
		BuyerClientID:   draft.BuyerClientID,
		SellerCompanyID: draft.SellerCompanyID,
		SellerCompany:   seller.Name,
		StatusInvoice:   draft.StatusInvoice,
		Source:          draft.Source,
		Subtotal:        draft.Subtotal,
		Taxes:           draft.Taxes,
		TotalAmount:     draft.TotalAmount,
		Items:           make([]marketplacedto.MarketplaceCartItemResponse, 0, len(items)),
	}

	for _, item := range items {
		res.Items = append(res.Items, marketplacedto.MarketplaceCartItemResponse{
			ID:               item.ID,
			InvoiceID:        draft.ID,
			ProductID:        item.ProductID,
			ProductName:      item.ProductName,
			ProductImagePath: item.ProductImagePath,
			Quantity:         item.Quantity,
			Price:            item.Price,
			Subtotal:         item.Subtotal,
		})
	}

	return res, nil
}
