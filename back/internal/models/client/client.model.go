package models

import "time"

type Client struct {
	ID             int
	UserID         *int
	CompanyID      *int
	FirstName      string
	LastName       string
	NameComplete   *string
	TotalInvoices  int
	TotalPurchased float64
	Email          string
	Phone          string
	CreatedAt      time.Time
	Status         int
	DeletedAt      *time.Time
	CompanyName     *string
}
