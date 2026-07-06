package invoiceService

import (
	"context"
	"crm-system-sales/internal/core/dto"
	errorHandler "crm-system-sales/internal/core/error"
	tenantHelper "crm-system-sales/internal/core/tenant"
	transaction "crm-system-sales/internal/core/transaction"
	"crm-system-sales/internal/modules/client"
	"crm-system-sales/internal/modules/company"
	inventoryservice "crm-system-sales/internal/modules/inventory"
	invoiceAccess "crm-system-sales/internal/modules/invoice/access"
	"crm-system-sales/internal/modules/invoice/constants"
	invoiceConstants "crm-system-sales/internal/modules/invoice/constants"
	invoicedto "crm-system-sales/internal/modules/invoice/dto"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"
	invoiceRepository "crm-system-sales/internal/modules/invoice/repository"
	invoiceitemrepo "crm-system-sales/internal/modules/invoice_item"
	invoicepaymentmodel "crm-system-sales/internal/modules/invoice_payment/models"
	invoicepaymentrepo "crm-system-sales/internal/modules/invoice_payment/repository"
	invoicePaymentWorkflow "crm-system-sales/internal/services/invoice_workflow/service"

	"database/sql"
	"log"
	"net/http"
)

type InvoiceService interface {
	CreateDraft(ctx context.Context, req *invoicedto.CreateInvoiceRequest) (*invoicedto.InvoiceResponse, error)
	Pay(ctx context.Context, invoiceID int, req *invoicedto.PayInvoiceRequest) (*invoicedto.PayInvoiceResponse, error)
	GetByID(ctx context.Context, id int) (*invoicedto.GetInvoiceResponse, error)
	Cancel(
		ctx context.Context,
		invoiceID int,
	) error
	GetCompanyInvoices(
		ctx context.Context,
		req *invoicedto.GetCompanyInvoicesRequest,
	) (*invoicedto.GetCompanyInvoicesResponse, error)
}

type invoiceService struct {
	db                     *sql.DB
	Repo                   invoiceRepository.InvoiceRepository
	ClientRepo             client.ClientRepository
	CompanyRepo            company.CompanyRepository
	InvoiceItemRepo        invoiceitemrepo.InvoiceItemRepository
	InventoryService       inventoryservice.InventoryService
	InvoicePaymentRepo     invoicepaymentrepo.InvoicePaymentRepository
	InvoicePaymentWorkflow invoicePaymentWorkflow.PayInvoiceWorkflow
}

func NewInvoiceService(
	db *sql.DB,
	repo invoiceRepository.InvoiceRepository,
	clientRepo client.ClientRepository,
	companyRepo company.CompanyRepository,
	invoiceItemRepo invoiceitemrepo.InvoiceItemRepository,
	inventoryService inventoryservice.InventoryService,
	invoicePaymentRepo invoicepaymentrepo.InvoicePaymentRepository,
	invoiceWorkflow invoicePaymentWorkflow.PayInvoiceWorkflow,
) InvoiceService {
	return &invoiceService{
		db:                     db,
		Repo:                   repo,
		ClientRepo:             clientRepo,
		CompanyRepo:            companyRepo,
		InvoiceItemRepo:        invoiceItemRepo,
		InventoryService:       inventoryService,
		InvoicePaymentRepo:     invoicePaymentRepo,
		InvoicePaymentWorkflow: invoiceWorkflow,
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

func (s *invoiceService) Pay(
	ctx context.Context,
	invoiceID int,
	req *invoicedto.PayInvoiceRequest,
) (*invoicedto.PayInvoiceResponse, error) {

	var err error

	tenant := tenantHelper.GetTenant(ctx)

	invoice, err := s.Repo.GetByID(
		ctx,
		invoiceID,
	)
	if err != nil {
		return nil, err
	}

	buyer, err := s.ClientRepo.GetByID(
		ctx,
		invoice.BuyerClientID,
	)
	if err != nil {
		return nil, err
	}

	err = invoiceAccess.CanPayInvoice(
		tenant,
		invoice,
		buyer,
	)
	if err != nil {
		return nil, err
	}

	remaining := invoice.TotalAmount - invoice.PaidAmount

	if req.Amount > remaining {
		return nil, errorHandler.NewAppError(
			http.StatusBadRequest,
			"El monto supera el saldo pendiente",
		)
	}

	payment := &invoicepaymentmodel.InvoicePayment{
		InvoiceID:     invoice.ID,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		PaidByUserID:  tenant.UserID,
	}

	status := invoice.StatusInvoice
	newPaidAmount := invoice.PaidAmount + req.Amount
	remainingAmount := invoice.TotalAmount - newPaidAmount

	err = transaction.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {
		err = s.InvoicePaymentRepo.Create(
			ctx,
			tx,
			payment,
		)
		if err != nil {
			return err
		}

		err = s.Repo.SetPaidAmount(
			ctx,
			tx,
			invoice.ID,
			newPaidAmount,
		)
		if err != nil {
			return err
		}

		if newPaidAmount >= invoice.TotalAmount {

			err = s.Repo.UpdateStatus(
				ctx,
				tx,
				invoice.ID,
				constants.InvoicePaid,
			)
			if err != nil {
				return err
			}

			err = s.InvoicePaymentWorkflow.FinalizePayment(
				ctx,
				tx,
				invoice.ID,
			)
			if err != nil {
				return err
			}

			status = constants.InvoicePaid
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	response := &invoicedto.PayInvoiceResponse{
		ID:              payment.ID,
		StatusInvoice:   status,
		TotalAmount:     invoice.TotalAmount,
		PaidAmount:      newPaidAmount,
		RemainingAmount: remainingAmount,
	}

	return response, nil
}

func (s *invoiceService) GetByID(
	ctx context.Context,
	id int,
) (*invoicedto.GetInvoiceResponse, error) {
	var err error

	tenant := tenantHelper.GetTenant(ctx)

	invoice, err := s.Repo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		return nil, err
	}

	buyer, err := s.ClientRepo.GetByID(
		ctx,
		invoice.BuyerClientID,
	)
	if err != nil {
		return nil, err
	}

	err = invoiceAccess.CanViewInvoice(
		tenant,
		invoice,
		buyer,
	)

	if err != nil {
		return nil, err
	}

	seller, err := s.CompanyRepo.GetByID(
		ctx,
		invoice.SellerCompanyID,
	)
	if err != nil {
		return nil, err
	}

	return &invoicedto.GetInvoiceResponse{
		ID:              invoice.ID,
		BuyerClientID:   invoice.BuyerClientID,
		SellerCompanyID: invoice.SellerCompanyID,
		CreatedByUserID: invoice.CreatedByUserID,
		StatusInvoice:   invoice.StatusInvoice,
		TotalAmount:     invoice.TotalAmount,
		PaidAmount:      invoice.PaidAmount,
		Status:          invoice.Status,
		Subtotal:        invoice.Subtotal,
		Taxes:           invoice.Taxes,
		CreatedAt:       invoice.CreatedAt,
		UpdatedAt:       invoice.UpdatedAt,
		DeletedAt:       invoice.DeletedAt,
		BuyerName:       buyer.FirstName + " " + buyer.LastName,
		SellerCompany:   seller.Name,
	}, nil
}

func (s *invoiceService) Cancel(
	ctx context.Context,
	invoiceID int,
) error {

	tenant := tenantHelper.GetTenant(ctx)

	invoice, err := s.Repo.GetByID(
		ctx,
		invoiceID,
	)
	if err != nil {
		return err
	}

	err = invoiceAccess.CanCancelInvoice(
		tenant,
		invoice,
	)
	if err != nil {
		return err
	}

	return transaction.RunInTransaction(
		ctx,
		s.db,
		func(tx *sql.Tx) error {

			if invoice.StatusInvoice == constants.InvoiceDraft ||
				invoice.StatusInvoice == constants.InvoicePending {

				items, err := s.InvoiceItemRepo.ListByInvoiceID(
					ctx,
					tx,
					invoiceID,
				)
				if err != nil {
					return err
				}

				for _, item := range items {

					err = s.InventoryService.ReleaseStock(
						ctx,
						tx,
						item.ProductID,
						item.Quantity,
					)
					if err != nil {
						return err
					}
				}
			}

			err := s.Repo.UpdateStatus(
				ctx,
				tx,
				invoice.ID,
				constants.InvoiceCanceled,
			)
			if err != nil {
				return err
			}

			return nil
		},
	)
}

func (s *invoiceService) GetCompanyInvoices(
	ctx context.Context,
	req *invoicedto.GetCompanyInvoicesRequest,
) (*invoicedto.GetCompanyInvoicesResponse, error) {
	tenant := tenantHelper.GetTenant(ctx)

	if tenant.CompanyID == nil {
		return nil, errorHandler.NewAppError(
			http.StatusForbidden,
			"Solo empresas",
		)
	}

	invoices, total, err := s.Repo.ListBySellerCompanyID(
		ctx,
		*tenant.CompanyID,
		req,
	)
	if err != nil {
		return nil, err
	}

	// Mapear a DTO de respuesta
	items := make([]invoicedto.CompanyInvoiceResponse, 0, len(invoices))
	for _, inv := range invoices {
		items = append(items, invoicedto.CompanyInvoiceResponse{
			ID:              inv.ID,
			BuyerClientID:   inv.BuyerClientID,
			SellerCompanyID: inv.SellerCompanyID,
			StatusInvoice:   inv.StatusInvoice,
			TotalAmount:     inv.TotalAmount,
			PaidAmount:      inv.PaidAmount,
			CreatedAt:       inv.CreatedAt,
			BuyerName:       inv.BuyerName,
    		SellerCompany:   inv.SellerName,
		})
	}

	// Meta info
	meta := dto.Meta{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &invoicedto.GetCompanyInvoicesResponse{
		Items: items,
		Meta:  meta,
	}, nil
}
