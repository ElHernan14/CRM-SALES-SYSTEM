package invoiceWorkflow

import (
	"context"
	errorHandler "crm-system-sales/internal/core/error"
	tenantHelper "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/core/transaction"
	"crm-system-sales/internal/core/utils"
	inventoryService "crm-system-sales/internal/modules/inventory"
	invoiceAccess "crm-system-sales/internal/modules/invoice/access"
	"crm-system-sales/internal/modules/invoice/constants"
	invoiceRepository "crm-system-sales/internal/modules/invoice/repository"
	invoiceItemRepository "crm-system-sales/internal/modules/invoice_item"
	"database/sql"
	"log"
	"net/http"
)

type SubmitInvoiceWorkflow interface {
	Submit(
		ctx context.Context,
		invoiceID int,
	) error
}

type submitInvoiceWorkflow struct {
	invoiceRepo      invoiceRepository.InvoiceRepository
	invoiceItemRepo  invoiceItemRepository.InvoiceItemRepository
	inventoryService inventoryService.InventoryService
	db               *sql.DB
}

func NewSubmitInvoiceWorkflow(
	invoiceRepo invoiceRepository.InvoiceRepository,
	invoiceItemRepo invoiceItemRepository.InvoiceItemRepository,
	inventoryService inventoryService.InventoryService,
	db *sql.DB,
) SubmitInvoiceWorkflow {
	return &submitInvoiceWorkflow{
		invoiceRepo:      invoiceRepo,
		invoiceItemRepo:  invoiceItemRepo,
		inventoryService: inventoryService,
		db:               db,
	}
}

func (s *submitInvoiceWorkflow) Submit(
	ctx context.Context,
	invoiceID int,
) error {

	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE SubmitInvoiceWorkflowService")(err)
	}()

	tenant := tenantHelper.GetTenant(ctx)

	invoice, err := s.invoiceRepo.GetByID(ctx, invoiceID)
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
	count, err := s.invoiceItemRepo.CountByInvoice(
		ctx,
		invoice.ID,
	)

	if err != nil {
		return errorHandler.NewAppError(
			http.StatusInternalServerError,
			"Error validando items de la invoice",
		)
	}

	if count == 0 {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"la invoice no posee items",
		)
	}

	//  validate subtotal
	if invoice.Subtotal <= 0 {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"subtotal inválido",
		)
	}

	err = transaction.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {

		return s.invoiceRepo.UpdateStatus(
			ctx,
			tx,
			invoice.ID,
			constants.InvoicePending,
		)
	})

	return err
}
