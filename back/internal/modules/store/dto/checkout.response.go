package dto

type CheckoutResponse struct {
	InvoiceID int    `json:"invoice_id"`
	Status    string `json:"status"`
}
