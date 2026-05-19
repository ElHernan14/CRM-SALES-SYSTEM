package invoiceitem

import (
	"context"
	"database/sql"

	invoiceItemModel "crm-system-sales/internal/modules/invoice_item/models"
)

type InvoiceItemRepository interface {
	Create(
		ctx context.Context,
		tx *sql.Tx,
		item *invoiceItemModel.InvoiceItem,
	) error
}

type invoiceItemRepository struct {
	db *sql.DB
}

func NewInvoiceItemRepository(db *sql.DB) InvoiceItemRepository {
	return &invoiceItemRepository{db: db}
}

func (r *invoiceItemRepository) Create(
	ctx context.Context,
	tx *sql.Tx,
	item *invoiceItemModel.InvoiceItem,
) error {

	query := `
		INSERT INTO invoice_item (
			invoice_id,
			product_id,
			product_name,
			quantity,
			price,
			subtotal
		)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id
	`

	return tx.QueryRowContext(
		ctx,
		query,
		item.InvoiceID,
		item.ProductID,
		item.ProductName,
		item.Quantity,
		item.Price,
		item.Subtotal,
	).Scan(&item.ID)
}
