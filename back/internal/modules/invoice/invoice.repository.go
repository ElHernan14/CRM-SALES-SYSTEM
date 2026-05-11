package invoice

import (
	"context"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"
	"database/sql"
)

type InvoiceRepository interface {
	Create(ctx context.Context, tx *sql.Tx, invoice *invoiceModel.Invoice) error
}

type invoiceRepository struct {
	DB *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *invoiceRepository {
	return &invoiceRepository{DB: db}
}

func (r *invoiceRepository) Create(
	ctx context.Context,
	tx *sql.Tx,
	invoice *invoiceModel.Invoice,
) error {

	query := `
		INSERT INTO invoice (
			buyer_client_id,
			seller_company_id,
			created_by_user_id,
			total_amount,
			status_invoice
		)
		VALUES ($1,$2,$3,$4,$5)
		RETURNING id, created_at
	`

	return tx.QueryRowContext(
		ctx,
		query,
		invoice.BuyerClientID,
		invoice.SellerCompanyID,
		invoice.CreatedByUserID,
		invoice.TotalAmount,
		invoice.StatusInvoice,
	).Scan(
		&invoice.ID,
		&invoice.CreatedAt,
	)
}
