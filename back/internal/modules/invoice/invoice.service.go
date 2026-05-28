package invoice

import (
	"context"
	errorHandler "crm-system-sales/internal/core/error"
	tenantHelper "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/core/transaction"
	"crm-system-sales/internal/core/utils"
	"crm-system-sales/internal/modules/client"
	"crm-system-sales/internal/modules/company"
	invoiceAccess "crm-system-sales/internal/modules/invoice/access"
	"crm-system-sales/internal/modules/invoice/constants"
	invoiceConstants "crm-system-sales/internal/modules/invoice/constants"
	invoicedto "crm-system-sales/internal/modules/invoice/dto"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"

	// invoiceItemRepo "crm-system-sales/internal/modules/invoice_item"
	"database/sql"
	"log"
	"net/http"
)

type InvoiceService interface {
	CreateDraft(ctx context.Context, req *invoicedto.CreateInvoiceRequest) (*invoicedto.InvoiceResponse, error)
	Submit(
		ctx context.Context,
		invoiceID int,
	) error
}

type invoiceService struct {
	db          *sql.DB
	Repo        InvoiceRepository
	ClientRepo  client.ClientRepository
	CompanyRepo company.CompanyRepository
	// InvoiceItemRepo invoiceItemRepo.InvoiceItemRepository
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
		log.Println("error fetching buyer client:", err)
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Cliente comprador no encontrado",
		)
	}

	//  empresa seller existe
	company, err := s.CompanyRepo.GetByID(ctx, req.SellerCompanyID)
	if err != nil {
		log.Println("error fetching seller company:", err)
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Empresa vendedora no encontrada",
		)
	}

	existingDraft, err := s.Repo.GetActiveDraft(
		ctx,
		req.BuyerClientID,
		req.SellerCompanyID,
	)

	if err != nil {
		log.Println("error checking existing draft:", err)
		return nil, errorHandler.NewAppError(
			http.StatusNotFound,
			"Error verificando draft existente",
		)
	}

	if existingDraft != nil {
		return &invoicedto.InvoiceResponse{
			ID:              existingDraft.ID,
			BuyerClientID:   existingDraft.BuyerClientID,
			SellerCompanyID: existingDraft.SellerCompanyID,
			StatusInvoice:   existingDraft.StatusInvoice,
			TotalAmount:     existingDraft.TotalAmount,
			CreatedAt:       existingDraft.CreatedAt,
			CreatedByUserID: existingDraft.CreatedByUserID,
		}, nil
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

func (s *invoiceService) Submit(
	ctx context.Context,
	invoiceID int,
) error {

	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE SubmitInvoice")(err)
	}()

	tenant := tenantHelper.GetTenant(ctx)

	invoice, err := s.Repo.GetByID(ctx, invoiceID)
	if err != nil {
		log.Println("error fetching invoice:", err)
		return errorHandler.NewAppError(
			http.StatusNotFound,
			"Invoice no encontrada",
		)
	}

	//  ownership
	err = invoiceAccess.CanEditDraftInvoice(
		tenant,
		invoice,
	)

	if err != nil {
		return err
	}

	//  status validation
	if invoice.StatusInvoice != constants.InvoiceDraft {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"solo invoices draft pueden enviarse",
		)
	}

	//  validate items
	// count, err := s.InvoiceItemRepo.CountByInvoice(
	// 	ctx,
	// 	invoice.ID,
	// )

	// if err != nil {
	// 	return errorHandler.NewAppError(
	// 		http.StatusInternalServerError,
	// 		"Error validando items de la invoice",
	// 	)
	// }

	// if count == 0 {
	// 	return errorHandler.NewAppError(
	// 		http.StatusBadRequest,
	// 		"la invoice no posee items",
	// 	)
	// }

	//  validate subtotal
	if invoice.Subtotal <= 0 {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"subtotal inválido",
		)
	}

	err = transaction.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {

		return s.Repo.UpdateStatus(
			ctx,
			tx,
			invoice.ID,
			constants.InvoicePending,
		)
	})

	return err
}
