package invoice

import (
	"context"
	errorHandler "crm-system-sales/internal/core/error"
	tenantHelper "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/modules/client"
	"crm-system-sales/internal/modules/company"
	invoiceAccess "crm-system-sales/internal/modules/invoice/access"
	invoiceConstants "crm-system-sales/internal/modules/invoice/constants"
	invoicedto "crm-system-sales/internal/modules/invoice/dto"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"
	"database/sql"
	"log"
	"net/http"
)

type InvoiceService interface {
	CreateDraft(ctx context.Context, req *invoicedto.CreateInvoiceRequest) (*invoicedto.InvoiceResponse, error)
}

type invoiceService struct {
	db          *sql.DB
	Repo        InvoiceRepository
	ClientRepo  client.ClientRepository
	CompanyRepo company.CompanyRepository
}

func NewInvoiceService(
	db *sql.DB,
	repo InvoiceRepository,
	clientRepo client.ClientRepository,
	companyRepo company.CompanyRepository,
) InvoiceService {
	return &invoiceService{
		db:          db,
		Repo:        repo,
		ClientRepo:  clientRepo,
		CompanyRepo: companyRepo,
	}
}

func (s *invoiceService) CreateDraft(
	ctx context.Context,
	req *invoicedto.CreateInvoiceRequest,
) (*invoicedto.InvoiceResponse, error) {

	tenant := tenantHelper.GetTenant(ctx)
	if tenant == nil {
		return nil, errorHandler.NewAppError(
			http.StatusUnauthorized,
			"Usuario no autenticado",
		)
	}

	//  buyer existe
	buyer, err := s.ClientRepo.GetByID(ctx, req.BuyerClientID)
	if err != nil {
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Cliente comprador no encontrado",
		)
	}

	//  empresa seller existe
	company, err := s.CompanyRepo.GetByID(ctx, req.SellerCompanyID)
	if err != nil {
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Empresa vendedora no encontrada",
		)
	}

	_ = company

	//  no comprarte a vos mismo
	if buyer.CompanyID != nil &&
		req.SellerCompanyID == *buyer.CompanyID {

		return nil, errorHandler.NewAppError(
			http.StatusBadRequest,
			"una empresa no puede facturarse a sí misma",
		)
	}

	//  validar access inicial
	err = invoiceAccess.CanCreateInvoice(
		tenant,
		buyer,
		req.SellerCompanyID,
	)

	if err != nil {
		return nil, err
	}

	invoice := &invoiceModel.Invoice{
		BuyerClientID:   req.BuyerClientID,
		SellerCompanyID: req.SellerCompanyID,
		CreatedByUserID: tenant.UserID,
		StatusInvoice:   invoiceConstants.InvoiceDraft,
		TotalAmount:     0,
	}

	err = s.Repo.Create(ctx, nil, invoice)
	if err != nil {
		log.Println("error creating invoice:", err)

		return nil, errorHandler.NewAppError(
			http.StatusInternalServerError,
			"Error creando invoice",
		)
	}

	return &invoicedto.InvoiceResponse{
		ID:              invoice.ID,
		BuyerClientID:   invoice.BuyerClientID,
		SellerCompanyID: invoice.SellerCompanyID,
		CreatedByUserID: invoice.CreatedByUserID,
		StatusInvoice:   invoice.StatusInvoice,
		TotalAmount:     invoice.TotalAmount,
		CreatedAt:       invoice.CreatedAt,
	}, nil
}
