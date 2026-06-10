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

type GetInvoiceResponse struct {
	ID              int `json:"id"`
	BuyerClientID   int `json:"buyer_client_id"`
	SellerCompanyID int `json:"seller_company_id"`
	CreatedByUserID int `json:"created_by_user_id"`

	StatusInvoice string `json:"status_invoice"`

	Subtotal    float64 `json:"subtotal"`
	Taxes       float64 `json:"taxes"`
	TotalAmount float64 `json:"total_amount"`
	PaidAmount  float64 `json:"paid_amount"`
	Status      int     `json:"status"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
