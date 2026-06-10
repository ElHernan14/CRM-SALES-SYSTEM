package dto

import (
	dto "crm-system-sales/internal/core/dto"
	"time"
)

type PaymentResponse struct {
	ID int `json:"id"`

	InvoiceID int `json:"invoice_id"`

	Amount float64 `json:"amount"`

	PaymentMethod string `json:"payment_method"`

	PaidByUserID int `json:"paid_by_user_id"`

	CreatedAt time.Time `json:"created_at"`
}

type GetInvoicePaymentsResponse struct {
	Items []PaymentResponse `json:"items"`

	Meta dto.Meta `json:"meta"`
}
