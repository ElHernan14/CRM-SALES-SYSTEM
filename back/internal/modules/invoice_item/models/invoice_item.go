package model

type InvoiceItem struct {
	ID        int
	InvoiceID int

	ProductID        int
	ProductName      string
	ProductImagePath *string

	Quantity int

	Price    float64
	Subtotal float64
}
