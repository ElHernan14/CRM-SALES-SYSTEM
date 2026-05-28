package invoiceitem

import (
	"context"
	"database/sql"

	"crm-system-sales/internal/core/utils"
	invoiceItemModel "crm-system-sales/internal/modules/invoice_item/models"
)

type InvoiceItemRepository interface {
	Create(ctx context.Context, tx *sql.Tx, item *invoiceItemModel.InvoiceItem) error
	GetByID(ctx context.Context, itemID int) (*invoiceItemModel.InvoiceItem, error)
	Update(ctx context.Context, tx *sql.Tx, item *invoiceItemModel.InvoiceItem) error
	Delete(ctx context.Context, tx *sql.Tx, itemID int) error
	GetByInvoiceAndProduct(ctx context.Context, invoiceID int, productID int) (*invoiceItemModel.InvoiceItem, error)
	CountByInvoice(
		ctx context.Context,
		invoiceID int,
	) (int, error)
	GetByInvoiceID(
		ctx context.Context,
		invoiceID int,
		page int,
		limit int,
	) ([]invoiceItemModel.InvoiceItem, int, error)
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
		AND status = 1
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
			updated_at = NOW(),
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

func (r *invoiceItemRepository) Delete(
	ctx context.Context,
	tx *sql.Tx,
	itemID int,
) error {

	query := `
		UPDATE invoice_item
		SET deleted_at = NOW(),
		    status = 0
		WHERE id = $1
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		itemID,
	)

	return err
}

func (r *invoiceItemRepository) GetByInvoiceAndProduct(
	ctx context.Context,
	invoiceID int,
	productID int,
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
		WHERE
			invoice_id = $1
			AND product_id = $2
			AND status = 1
	`

	item := &invoiceItemModel.InvoiceItem{}

	err := r.db.QueryRowContext(
		ctx,
		query,
		invoiceID,
		productID,
	).Scan(
		&item.ID,
		&item.InvoiceID,
		&item.ProductID,
		&item.ProductName,
		&item.Quantity,
		&item.Price,
		&item.Subtotal,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *invoiceItemRepository) CountByInvoice(
	ctx context.Context,
	invoiceID int,
) (int, error) {

	query := `
		SELECT COUNT(*)
		FROM invoice_item
		WHERE
			invoice_id = $1
			AND status = 1
	`

	var count int

	err := r.db.QueryRowContext(
		ctx,
		query,
		invoiceID,
	).Scan(&count)

	return count, err
}

func (r *invoiceItemRepository) GetByInvoiceID(
	ctx context.Context,
	invoiceID int,
	page int,
	limit int,
) ([]invoiceItemModel.InvoiceItem, int, error) {
	var err error
	defer func() {
		utils.Trace(ctx, "REPO GetByInvoiceID")(err)
	}()

	baseQuery := `
			 FROM invoice_item
			WHERE invoice_id = $1
			AND status = 1 `

	countQuery := `SELECT COUNT(*) ` + baseQuery
	var total int
	err = r.db.QueryRowContext(ctx, countQuery, invoiceID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	dataQuery := `SELECT id, product_id, product_name, quantity, price, subtotal 
				` + baseQuery + `
				ORDER BY id ASC
				LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, dataQuery, invoiceID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := []invoiceItemModel.InvoiceItem{}

	for rows.Next() {
		var item invoiceItemModel.InvoiceItem
		err := rows.Scan(
			&item.ID,
			&item.ProductID,
			&item.ProductName,
			&item.Quantity,
			&item.Price,
			&item.Subtotal,
		)
		if err != nil {
			return nil, 0, err
		}
		items = append(items, item)
	}

	return items, total, nil
}
