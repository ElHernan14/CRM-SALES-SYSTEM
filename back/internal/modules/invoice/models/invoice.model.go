package model

import "time"

type Invoice struct {
	ID              int
	BuyerClientID   int
	SellerCompanyID int
	CreatedByUserID int

	StatusInvoice string

	Subtotal    float64
	Taxes       float64
	TotalAmount float64
	PaidAmount  float64

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time

	Status int

	BuyerName  string
	SellerName string
	ItemCount  int
}
