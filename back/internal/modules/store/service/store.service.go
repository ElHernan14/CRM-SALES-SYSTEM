package service

import (
	"context"
	"database/sql"
	"net/http"

	metadto "crm-system-sales/internal/core/dto"
	errorHandler "crm-system-sales/internal/core/error"
	tenantHelper "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/modules/company"
	"crm-system-sales/internal/modules/invoice/constants"
	invoicemodel "crm-system-sales/internal/modules/invoice/models"
	invoicerepo "crm-system-sales/internal/modules/invoice/repository"
	invoiceitemrepo "crm-system-sales/internal/modules/invoice_item"
	productrepo "crm-system-sales/internal/modules/product"
	storedto "crm-system-sales/internal/modules/store/dto"
	submitInvoiceWorkflow "crm-system-sales/internal/services/invoice_workflow/service"
)

type StoreService interface {
	GetProducts(ctx context.Context, req *storedto.GetStoreProductsRequest) (*storedto.GetStoreProductsResponse, error)
	GetProductByID(ctx context.Context, id int) (*storedto.StoreProductDetailResponse, error)
	GetCart(ctx context.Context, sellerCompanyID int) (*storedto.CartResponse, error)
	EnsureCart(ctx context.Context, req *storedto.EnsureCartRequest) (*storedto.EnsureCartResponse, error)
	Checkout(ctx context.Context, req *storedto.CheckoutRequest) (*storedto.CheckoutResponse, error)
}

type storeService struct {
	db              *sql.DB
	ProductRepo     productrepo.ProductRepository
	InvoiceRepo     invoicerepo.InvoiceRepository
	InvoiceItemRepo invoiceitemrepo.InvoiceItemRepository
	CompanyRepo     company.CompanyRepository
	SubmitWorkflow  submitInvoiceWorkflow.SubmitInvoiceWorkflow
}

func NewStoreService(
	db *sql.DB,
	productRepo productrepo.ProductRepository,
	invoiceRepo invoicerepo.InvoiceRepository,
	invoiceItemRepo invoiceitemrepo.InvoiceItemRepository,
	companyRepo company.CompanyRepository,
	submitWorkflow submitInvoiceWorkflow.SubmitInvoiceWorkflow,
) StoreService {
	return &storeService{db: db, ProductRepo: productRepo, InvoiceRepo: invoiceRepo, InvoiceItemRepo: invoiceItemRepo, CompanyRepo: companyRepo, SubmitWorkflow: submitWorkflow}
}

func (s *storeService) GetProducts(ctx context.Context, req *storedto.GetStoreProductsRequest) (*storedto.GetStoreProductsResponse, error) {
	products, total, err := s.ProductRepo.ListAvailableProducts(ctx, req)
	if err != nil {
		return nil, err
	}

	items := make([]storedto.StoreProductResponse, 0, len(products))
	for _, p := range products {
		available := p.Stock - p.ReservedStock
		items = append(items, storedto.StoreProductResponse{ID: p.ID, CompanyID: p.CompanyID, CompanyName: p.CompanyName, Name: p.Name, Kind: p.Kind, CategoryID: p.CategoryID, Category: p.Category, TypeID: p.TypeID, Type: p.Type, Description: p.Description, Price: p.Price, AvailableStock: available, ImagePath: p.ImagePath})
	}

	meta := metadto.NewMeta(req.Page, req.Limit, total)
	return &storedto.GetStoreProductsResponse{Items: items, Meta: meta}, nil
}

func (s *storeService) GetProductByID(ctx context.Context, id int) (*storedto.StoreProductDetailResponse, error) {
	if id <= 0 {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "product_id invalido")
	}

	product, err := s.ProductRepo.GetAvailableProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errorHandler.NewAppError(http.StatusNotFound, "Producto no encontrado")
	}

	available := product.Stock - product.ReservedStock
	return &storedto.StoreProductDetailResponse{
		ID:             product.ID,
		CompanyID:      product.CompanyID,
		CompanyName:    product.CompanyName,
		Name:           product.Name,
		Description:    product.Description,
		Kind:           product.Kind,
		CategoryID:     product.CategoryID,
		Category:       product.Category,
		TypeID:         product.TypeID,
		Type:           product.Type,
		Price:          product.Price,
		AvailableStock: available,
		ImagePath:      product.ImagePath,
	}, nil
}

func (s *storeService) GetCart(ctx context.Context, sellerCompanyID int) (*storedto.CartResponse, error) {
	tenant := tenantHelper.GetTenant(ctx)
	if tenant == nil || tenant.ClientID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "Solo clientes compradores")
	}

	if err := s.validateSeller(ctx, tenant, sellerCompanyID); err != nil {
		return nil, err
	}

	draft, err := s.InvoiceRepo.GetActiveDraft(ctx, *tenant.ClientID, sellerCompanyID)
	if err != nil {
		return nil, err
	}
	if draft == nil {
		return nil, nil
	}

	return s.buildCartResponse(ctx, draft)
}

func (s *storeService) EnsureCart(ctx context.Context, req *storedto.EnsureCartRequest) (*storedto.EnsureCartResponse, error) {
	tenant := tenantHelper.GetTenant(ctx)
	if tenant == nil || tenant.ClientID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "Solo clientes compradores")
	}

	if err := s.validateSeller(ctx, tenant, req.SellerCompanyID); err != nil {
		return nil, err
	}

	draft, err := s.InvoiceRepo.GetActiveDraft(ctx, *tenant.ClientID, req.SellerCompanyID)
	if err != nil {
		return nil, err
	}
	if draft == nil {
		draft = &invoicemodel.Invoice{BuyerClientID: *tenant.ClientID, SellerCompanyID: req.SellerCompanyID, CreatedByUserID: tenant.UserID, StatusInvoice: constants.InvoiceDraft, TotalAmount: 0}
		if err := s.InvoiceRepo.Create(ctx, nil, draft); err != nil {
			return nil, errorHandler.NewAppError(http.StatusInternalServerError, "Error creando carrito")
		}
	}

	return &storedto.EnsureCartResponse{InvoiceID: draft.ID, BuyerClientID: draft.BuyerClientID, SellerCompanyID: draft.SellerCompanyID, StatusInvoice: draft.StatusInvoice}, nil
}

func (s *storeService) Checkout(ctx context.Context, req *storedto.CheckoutRequest) (*storedto.CheckoutResponse, error) {
	tenant := tenantHelper.GetTenant(ctx)
	if tenant == nil {
		return nil, errorHandler.NewAppError(http.StatusUnauthorized, "Usuario no autenticado")
	}

	var invoiceID int
	if req != nil && req.InvoiceID != nil {
		invoiceID = *req.InvoiceID
	} else if req != nil && req.SellerCompanyID != nil {
		if tenant.ClientID == nil {
			return nil, errorHandler.NewAppError(http.StatusForbidden, "Solo clientes compradores")
		}
		draft, err := s.InvoiceRepo.GetActiveDraft(ctx, *tenant.ClientID, *req.SellerCompanyID)
		if err != nil {
			return nil, err
		}
		if draft == nil {
			return nil, errorHandler.NewAppError(http.StatusNotFound, "Carrito no encontrado")
		}
		invoiceID = draft.ID
	} else {
		draft, err := s.InvoiceRepo.GetActiveDraftByTenant(ctx, tenant)
		if err != nil {
			return nil, err
		}
		if draft == nil {
			return nil, errorHandler.NewAppError(http.StatusNotFound, "Carrito no encontrado")
		}
		invoiceID = draft.ID
	}

	if err := s.SubmitWorkflow.Submit(ctx, invoiceID); err != nil {
		return nil, err
	}

	return &storedto.CheckoutResponse{InvoiceID: invoiceID, Status: constants.InvoicePending}, nil
}

func (s *storeService) validateSeller(ctx context.Context, tenant *tenantHelper.TenantContext, sellerCompanyID int) error {
	if sellerCompanyID <= 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "seller_company_id invalido")
	}
	if tenant.CompanyID != nil && *tenant.CompanyID == sellerCompanyID {
		return errorHandler.NewAppError(http.StatusBadRequest, "una empresa no puede comprarse a si misma")
	}
	seller, err := s.CompanyRepo.GetByID(ctx, sellerCompanyID)
	if err != nil {
		return err
	}
	if seller == nil {
		return errorHandler.NewAppError(http.StatusNotFound, "Empresa proveedora no encontrada")
	}
	return nil
}

func (s *storeService) buildCartResponse(ctx context.Context, draft *invoicemodel.Invoice) (*storedto.CartResponse, error) {
	seller, err := s.CompanyRepo.GetByID(ctx, draft.SellerCompanyID)
	if err != nil {
		return nil, err
	}
	if seller == nil {
		return nil, errorHandler.NewAppError(http.StatusNotFound, "Empresa proveedora no encontrada")
	}

	items, _, err := s.InvoiceItemRepo.GetByInvoiceID(ctx, draft.ID, 1, 100)
	if err != nil {
		return nil, err
	}

	res := &storedto.CartResponse{InvoiceID: draft.ID, BuyerClientID: draft.BuyerClientID, SellerCompanyID: draft.SellerCompanyID, SellerCompany: seller.Name, StatusInvoice: draft.StatusInvoice, Subtotal: draft.Subtotal, Taxes: draft.Taxes, TotalAmount: draft.TotalAmount, Items: make([]storedto.CartItemResponse, 0, len(items))}
	for _, item := range items {
		res.Items = append(res.Items, storedto.CartItemResponse{ID: item.ID, InvoiceID: draft.ID, ProductID: item.ProductID, ProductName: item.ProductName, Quantity: item.Quantity, Price: item.Price, Subtotal: item.Subtotal})
	}
	return res, nil
}
