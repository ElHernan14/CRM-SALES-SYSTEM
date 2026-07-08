package invoicepaymentmodel

import "time"

type InvoicePayment struct {
	ID            int       `db:"id" json:"id"`
	InvoiceID     int       `db:"invoice_id" json:"invoice_id"`
	Amount        float64   `db:"amount" json:"amount"`
	PaymentMethod string    `db:"payment_method" json:"payment_method"`
	PaidByUserID  int       `db:"paid_by_user_id" json:"paid_by_user_id"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`

	ClientName string `db:"client_name" json:"client_name"`
}
