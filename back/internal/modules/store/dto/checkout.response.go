package dto

type CheckoutResponse struct {
	InvoiceID int    `json:"invoice_id" example:"15"`
	Status    string `json:"status" example:"pending"`
}
