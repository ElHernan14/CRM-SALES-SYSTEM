package invoicedto

type PayInvoiceResponse struct {
	ID              int     `json:"payment_id"`
	StatusInvoice   string  `json:"status_invoice"`
	TotalAmount     float64 `json:"total_amount"`
	PaidAmount      float64 `json:"paid_amount"`
	RemainingAmount float64 `json:"remaining_amount"`
}
