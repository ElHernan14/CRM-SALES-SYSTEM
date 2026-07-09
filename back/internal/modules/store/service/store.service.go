package service

import (
	"context"
	metadto "crm-system-sales/internal/core/dto"
	errorHandler "crm-system-sales/internal/core/error"
	tenantHelper "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/modules/invoice/constants"
	invoicerepo "crm-system-sales/internal/modules/invoice/repository"
	productrepo "crm-system-sales/internal/modules/product"
	storedto "crm-system-sales/internal/modules/store/dto"
	submitInvoiceWorkflow "crm-system-sales/internal/services/invoice_workflow/service"
	"database/sql"
	"net/http"
)

type StoreService interface {
	GetProducts(ctx context.Context, req *storedto.GetStoreProductsRequest) (*storedto.GetStoreProductsResponse, error)
	Checkout(ctx context.Context) (*storedto.CheckoutResponse, error)
}

type storeService struct {
	db             *sql.DB
	ProductRepo    productrepo.ProductRepository
	InvoiceRepo    invoicerepo.InvoiceRepository
	SubmitWorkflow submitInvoiceWorkflow.SubmitInvoiceWorkflow
}

func NewStoreService(db *sql.DB, repo productrepo.ProductRepository, submitWorkflow submitInvoiceWorkflow.SubmitInvoiceWorkflow) StoreService {
	return &storeService{db: db, ProductRepo: repo, SubmitWorkflow: submitWorkflow}
}

func (s *storeService) GetProducts(ctx context.Context, req *storedto.GetStoreProductsRequest) (*storedto.GetStoreProductsResponse, error) {
	products, total, err := s.ProductRepo.ListAvailableProducts(ctx, req)
	if err != nil {
		return nil, err
	}

	items := make([]storedto.StoreProductResponse, 0, len(products))
	for _, p := range products {
		available := p.Stock - p.ReservedStock
		items = append(items, storedto.StoreProductResponse{
			ID:             p.ID,
			CompanyID:      p.CompanyID,
			CompanyName:    p.CompanyName,
			Name:           p.Name,
			Type:           p.Type,
			Description:    p.Description,
			Price:          p.Price,
			AvailableStock: available,
		})
	}

	meta := metadto.Meta{Page: req.Page, Limit: req.Limit, Total: total}
	return &storedto.GetStoreProductsResponse{Items: items, Meta: meta}, nil
}

func (s *storeService) Checkout(ctx context.Context) (*storedto.CheckoutResponse, error) {
	tenant := tenantHelper.GetTenant(ctx)
	if tenant == nil {
		return nil, errorHandler.NewAppError(http.StatusUnauthorized, "Usuario no autenticado")
	}

	draft, err := s.InvoiceRepo.GetActiveDraftByTenant(ctx, tenant)
	if err != nil {
		return nil, err
	}

	err = s.SubmitWorkflow.Submit(ctx, draft.ID)
	if err != nil {
		return nil, err
	}

	return &storedto.CheckoutResponse{InvoiceID: draft.ID, Status: constants.InvoicePending}, nil
}
