package invoicedto

import "time"

type InvoiceResponse struct {
	ID              int       `json:"id"`
	BuyerClientID   int       `json:"buyer_client_id"`
	SellerCompanyID int       `json:"seller_company_id"`
	CreatedByUserID int       `json:"created_by_user_id"`
	StatusInvoice   string    `json:"status_invoice"`
	TotalAmount     float64   `json:"total_amount"`
	CreatedAt       time.Time `json:"created_at"`
}
