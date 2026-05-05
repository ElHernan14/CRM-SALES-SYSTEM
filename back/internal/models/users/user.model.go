package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	CompanyID    *int      `json:"company_id,omitempty"`
	Status       int       `json:"status"`
}
