package repository

import (
	"context"
	invoicepaymentmodel "crm-system-sales/internal/modules/invoice_payment/models"
	"database/sql"
)

type InvoicePaymentRepository interface {
	Create(
		ctx context.Context,
		tx *sql.Tx,
		payment *invoicepaymentmodel.InvoicePayment,
	) error
}

type invoicePaymentRepository struct {
	DB *sql.DB
}

func NewInvoicePaymentRepository(db *sql.DB) InvoicePaymentRepository {
	return &invoicePaymentRepository{DB: db}
}

func (r *invoicePaymentRepository) Create(
	ctx context.Context,
	tx *sql.Tx,
	payment *invoicepaymentmodel.InvoicePayment,
) error {

	query := `
		INSERT INTO invoice_payment (
			invoice_id,
			amount,
			payment_method,
			paid_by_user_id
		)
		VALUES (
			$1,$2,$3,$4
		)
		RETURNING id, created_at
	`

	return tx.QueryRowContext(
		ctx,
		query,
		payment.InvoiceID,
		payment.Amount,
		payment.PaymentMethod,
		payment.PaidByUserID,
	).Scan(
		&payment.ID,
		&payment.CreatedAt,
	)
}
