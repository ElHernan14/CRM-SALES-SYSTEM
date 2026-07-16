package dto

type CheckoutRequest struct {
	InvoiceID       *int `json:"invoice_id,omitempty" validate:"omitempty,gt=0"`
	SellerCompanyID *int `json:"seller_company_id,omitempty" validate:"omitempty,gt=0"`
}
