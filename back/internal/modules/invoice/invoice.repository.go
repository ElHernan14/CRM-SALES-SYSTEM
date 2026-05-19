package invoice

import (
	"context"
	"crm-system-sales/internal/core/utils"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"
	"database/sql"
)

type InvoiceRepository interface {
	Create(ctx context.Context, tx *sql.Tx, invoice *invoiceModel.Invoice) error
	GetByID(ctx context.Context, id int) (*invoiceModel.Invoice, error)
	RecalculateInvoiceTotals(ctx context.Context, tx *sql.Tx, invoiceID int) error
}

type invoiceRepository struct {
	DB *sql.DB
}

func NewInvoiceRepository(db *sql.DB) InvoiceRepository {
	return &invoiceRepository{DB: db}
}

const DefaultTaxRate = 0.21 // 21%

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
	if tx != nil {
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
	} else {
		return r.DB.QueryRowContext(
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
}

func (r *invoiceRepository) GetByID(ctx context.Context, id int) (*invoiceModel.Invoice, error) {
	var err error
	defer func() {
		utils.Trace(ctx, "REPO GetInvoiceByID")(err)
	}()

	query := `
        SELECT id, buyer_client_id, seller_company_id, created_by_user_id,
               total_amount, status_invoice, status
        FROM invoice
        WHERE id = $1
    `

	var inv invoiceModel.Invoice

	err = r.DB.QueryRowContext(ctx, query, id).Scan(
		&inv.ID,
		&inv.BuyerClientID,
		&inv.SellerCompanyID,
		&inv.CreatedByUserID,
		&inv.TotalAmount,
		&inv.StatusInvoice,
		&inv.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &inv, nil
}

func (r *invoiceRepository) RecalculateInvoiceTotals(
	ctx context.Context,
	tx *sql.Tx,
	invoiceID int,
) error {

	//  subtotal items
	var subtotal float64

	querySubtotal := `
		SELECT COALESCE(SUM(subtotal), 0)
		FROM invoice_item
		WHERE invoice_id = $1
	`

	err := tx.QueryRowContext(
		ctx,
		querySubtotal,
		invoiceID,
	).Scan(&subtotal)

	if err != nil {
		return err
	}

	//  taxes
	taxes := subtotal * DefaultTaxRate

	//  total
	total := subtotal + taxes

	//  update invoice
	queryUpdate := `
		UPDATE invoice
		SET
			subtotal = $1,
			taxes = $2,
			total_amount = $3,
			updated_at = NOW()
		WHERE id = $4
	`

	_, err = tx.ExecContext(
		ctx,
		queryUpdate,
		subtotal,
		taxes,
		total,
		invoiceID,
	)

	return err
}
