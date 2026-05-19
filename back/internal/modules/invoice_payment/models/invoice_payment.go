package model

import "time"

type InvoicePayment struct {
	ID int

	InvoiceID int
	Amount    float64

	PaymentMethod *string

	PaidByUserID int

	CreatedAt time.Time
}
