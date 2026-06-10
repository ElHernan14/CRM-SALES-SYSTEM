package dto

type GetInvoicePaymentsRequest struct {
	Page int `json:"page" validate:"gte=1"`

	Limit int `json:"limit" validate:"gte=1,lte=100"`

	PaymentMethod string `json:"payment_method" validate:"omitempty,payment_method"`

	SortColumn string `json:"sort_column" validate:"omitempty,oneof=created_at amount"`

	Order string `json:"order" validate:"omitempty,oneof=asc desc"`
}
