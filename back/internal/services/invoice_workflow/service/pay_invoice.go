package invoiceWorkflow

import (
	"context"
	inventoryservice "crm-system-sales/internal/modules/inventory"
	invoiceitemrepository "crm-system-sales/internal/modules/invoice_item"
	"database/sql"
)

type PayInvoiceWorkflow interface {
	FinalizePayment(ctx context.Context, tx *sql.Tx, invoiceID int) error
}

type payInvoiceWorkflow struct {
	invoiceItemRepo invoiceitemrepository.InvoiceItemRepository
	inventorySvc    inventoryservice.InventoryService
}

func NewPayInvoiceWorkflow(
	invoiceItemRepo invoiceitemrepository.InvoiceItemRepository,
	inventorySvc inventoryservice.InventoryService,
) PayInvoiceWorkflow {
	return &payInvoiceWorkflow{
		invoiceItemRepo: invoiceItemRepo,
		inventorySvc:    inventorySvc,
	}
}

func (s *payInvoiceWorkflow) FinalizePayment(
	ctx context.Context,
	tx *sql.Tx,
	invoiceID int,
) error {

	items, err := s.invoiceItemRepo.ListByInvoiceID(
		ctx,
		tx,
		invoiceID,
	)
	if err != nil {
		return err
	}

	for _, item := range items {

		err = s.inventorySvc.FinalizeReservedStock(
			ctx,
			tx,
			item.ProductID,
			item.Quantity,
		)
		if err != nil {
			return err
		}

	}

	return nil
}
