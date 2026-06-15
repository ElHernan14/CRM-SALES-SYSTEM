package invoiceWorkflow

import (
	"context"
	errorHandler "crm-system-sales/internal/core/error"
	tenantHelper "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/core/transaction"
	clientRepo "crm-system-sales/internal/modules/client"
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
	clientRepo       clientRepo.ClientRepository
	inventoryService inventoryService.InventoryService
	db               *sql.DB
}

func NewSubmitInvoiceWorkflow(
	invoiceRepo invoiceRepository.InvoiceRepository,
	invoiceItemRepo invoiceItemRepository.InvoiceItemRepository,
	clientRepo clientRepo.ClientRepository,
	inventoryService inventoryService.InventoryService,
	db *sql.DB,
) SubmitInvoiceWorkflow {
	return &submitInvoiceWorkflow{
		invoiceRepo:      invoiceRepo,
		invoiceItemRepo:  invoiceItemRepo,
		clientRepo:       clientRepo,
		inventoryService: inventoryService,
		db:               db,
	}
}

func (s *submitInvoiceWorkflow) Submit(
	ctx context.Context,
	invoiceID int,
) error {

	var err error

	tenant := tenantHelper.GetTenant(ctx)

	invoice, err := s.invoiceRepo.GetByID(ctx, invoiceID)
	if err != nil {
		log.Println("error fetching invoice:", err)
		return errorHandler.NewAppError(
			http.StatusNotFound,
			"Invoice no encontrada",
		)
	}

	//  buyer
	buyer, err := s.clientRepo.GetByID(ctx, invoice.BuyerClientID)
	if err != nil {
		return errorHandler.NewAppError(
			http.StatusNotFound,
			"Cliente no encontrado",
		)
	}

	//  ownership
	err = invoiceAccess.CanMutateDraftInvoice(
		tenant,
		invoice,
		buyer,
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

	log.Printf("Invoice %d tiene subtotal %.2f", invoice.ID, invoice.Subtotal)
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
