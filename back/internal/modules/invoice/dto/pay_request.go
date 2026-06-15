package invoicedto

type PayInvoiceRequest struct {
	Amount        float64 `json:"amount" validate:"required,gt=0" example:"1500"`
	PaymentMethod string  `json:"payment_method" validate:"required,payment_method,max=50" example:"cash"`
}
