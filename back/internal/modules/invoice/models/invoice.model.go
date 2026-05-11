package model

import "time"

type Invoice struct {
	ID              int       `db:"id"`
	BuyerClientID   int       `db:"buyer_client_id"`
	SellerCompanyID int       `db:"seller_company_id"`
	TotalAmount     float64   `db:"total_amount"`
	CreatedAt       time.Time `db:"created_at"`
	CreatedByUserID int       `db:"created_by_user_id"`
	StatusInvoice   string    `db:"status_invoice"`
}
