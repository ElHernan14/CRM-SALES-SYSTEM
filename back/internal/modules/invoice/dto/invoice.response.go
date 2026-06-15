package invoicedto

import "time"

type InvoiceResponse struct {
	ID              int       `json:"id" example:"15"`
	BuyerClientID   int       `json:"buyer_client_id" example:"5"`
	SellerCompanyID int       `json:"seller_company_id" example:"2"`
	CreatedByUserID int       `json:"created_by_user_id" example:"10"`
	StatusInvoice   string    `json:"status_invoice" example:"draft"`
	TotalAmount     float64   `json:"total_amount" example:"0"`
	CreatedAt       time.Time `json:"created_at" example:"2026-06-15T18:30:00Z"`
}

type GetInvoiceResponse struct {
	ID              int `json:"id" example:"15"`
	BuyerClientID   int `json:"buyer_client_id" example:"8"`
	SellerCompanyID int `json:"seller_company_id" example:"2"`
	CreatedByUserID int `json:"created_by_user_id" example:"12"`

	StatusInvoice string `json:"status_invoice" example:"pending"`

	Subtotal    float64 `json:"subtotal" example:"1000"`
	Taxes       float64 `json:"taxes" example:"210"`
	TotalAmount float64 `json:"total_amount" example:"1210"`
	PaidAmount  float64 `json:"paid_amount" example:"500"`
	Status      int     `json:"status" example:"1"`

	CreatedAt time.Time  `json:"created_at" example:"2026-06-15T10:00:00Z"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
