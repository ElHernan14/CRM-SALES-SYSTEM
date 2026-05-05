package models

import "time"

type Client struct {
	ID        int
	UserID    *int
	CompanyID *int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CreatedAt time.Time
	Status    int
	DeletedAt *time.Time
}
