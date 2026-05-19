package invoiceitem

import (
	"context"
	"database/sql"

	invoiceItemModel "crm-system-sales/internal/modules/invoice_item/models"
)

type InvoiceItemRepository interface {
	Create(ctx context.Context, tx *sql.Tx, item *invoiceItemModel.InvoiceItem) error
	GetByID(ctx context.Context, itemID int) (*invoiceItemModel.InvoiceItem, error)
	Update(ctx context.Context, tx *sql.Tx, item *invoiceItemModel.InvoiceItem) error
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

func (r *invoiceItemRepository) GetByID(
	ctx context.Context,
	itemID int,
) (*invoiceItemModel.InvoiceItem, error) {

	query := `
		SELECT
			id,
			invoice_id,
			product_id,
			product_name,
			quantity,
			price,
			subtotal
		FROM invoice_item
		WHERE id = $1
	`

	item := &invoiceItemModel.InvoiceItem{}

	err := r.db.QueryRowContext(
		ctx,
		query,
		itemID,
	).Scan(
		&item.ID,
		&item.InvoiceID,
		&item.ProductID,
		&item.ProductName,
		&item.Quantity,
		&item.Price,
		&item.Subtotal,
	)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *invoiceItemRepository) Update(
	ctx context.Context,
	tx *sql.Tx,
	item *invoiceItemModel.InvoiceItem,
) error {

	query := `
		UPDATE invoice_item
		SET
			quantity = $1,
			subtotal = $2
		WHERE id = $3
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		item.Quantity,
		item.Subtotal,
		item.ID,
	)

	return err
}
