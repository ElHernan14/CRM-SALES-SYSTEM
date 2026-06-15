package invoicedto

type PayInvoiceResponse struct {
	ID              int     `json:"payment_id" example:"21"`
	StatusInvoice   string  `json:"status_invoice" example:"paid"`
	TotalAmount     float64 `json:"total_amount" example:"1500"`
	PaidAmount      float64 `json:"paid_amount" example:"1500"`
	RemainingAmount float64 `json:"remaining_amount" example:"0"`
}
