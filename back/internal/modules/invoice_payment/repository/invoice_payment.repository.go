package repository

import (
	"context"
	invoicepaymentdto "crm-system-sales/internal/modules/invoice_payment/dto"
	invoicepaymentmodel "crm-system-sales/internal/modules/invoice_payment/models"
	"database/sql"
	"fmt"
	"strings"
)

type InvoicePaymentRepository interface {
	Create(
		ctx context.Context,
		tx *sql.Tx,
		payment *invoicepaymentmodel.InvoicePayment,
	) error
	ListByInvoiceID(
		ctx context.Context,
		invoiceID int,
		req *invoicepaymentdto.GetInvoicePaymentsRequest,
	) ([]*invoicepaymentmodel.InvoicePayment, int, error)
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

func (r *invoicePaymentRepository) ListByInvoiceID(
	ctx context.Context,
	invoiceID int,
	req *invoicepaymentdto.GetInvoicePaymentsRequest,
) ([]*invoicepaymentmodel.InvoicePayment, int, error) {
	var err error

	allowedSortColumns := map[string]string{
		"created_at": "created_at",
		"amount":     "amount",
	}

	allowedOrders := map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}

	selectQuery := `
		SELECT
			ip.id,
			ip.invoice_id,
			ip.amount,
			ip.payment_method,
			ip.paid_by_user_id,
			ip.created_at,
			COALESCE(c.first_name, '') AS client_first_name,
			COALESCE(c.last_name, '') AS client_last_name
			`
	baseQuery := `
		FROM invoice_payment ip
		LEFT JOIN client c ON c.user_id = ip.paid_by_user_id
		WHERE ip.invoice_id = $1 
	`

	args := []interface{}{
		invoiceID,
	}

	argPos := 2

	if req.PaymentMethod != "" {

		baseQuery += fmt.Sprintf(
			" AND payment_method = $%d",
			argPos,
		)

		args = append(
			args,
			req.PaymentMethod,
		)

		argPos++
	}

	countQuery := "SELECT COUNT(*) " + baseQuery

	var total int
	if err := r.DB.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	sortColumn := "created_at"
	order := "DESC"

	if v, ok := allowedSortColumns[req.SortColumn]; ok {
		sortColumn = v
	}

	if v, ok := allowedOrders[strings.ToLower(req.Order)]; ok {
		order = v
	}

	offset := (req.Page - 1) * req.Limit

	baseQuery += fmt.Sprintf(
		" ORDER BY %s %s",
		sortColumn,
		order,
	)

	baseQuery += fmt.Sprintf(
		" LIMIT $%d OFFSET $%d",
		argPos,
		argPos+1,
	)

	args = append(
		args,
		req.Limit,
		offset,
	)

	rows, err := r.DB.QueryContext(
		ctx,
		selectQuery+baseQuery,
		args...,
	)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	payments :=
		make(
			[]*invoicepaymentmodel.InvoicePayment,
			0,
		)

	for rows.Next() {

		var payment invoicepaymentmodel.InvoicePayment

		var clientFirstName, clientLastName sql.NullString

		err := rows.Scan(
			&payment.ID,
			&payment.InvoiceID,
			&payment.Amount,
			&payment.PaymentMethod,
			&payment.PaidByUserID,
			&payment.CreatedAt,
			&clientFirstName,
			&clientLastName,
		)
		if err != nil {
			return nil, 0, err
		}

		payment.ClientName = strings.TrimSpace(clientLastName.String + " " + clientFirstName.String)

		payments = append(
			payments,
			&payment,
		)
	}

	return payments, total, nil
}
