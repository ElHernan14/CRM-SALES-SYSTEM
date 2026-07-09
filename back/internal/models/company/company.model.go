package models

import "time"

type Company struct {
	ID          int        `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description *string    `json:"description,omitempty" db:"description"`
	LogoPath    *string    `json:"logo_path,omitempty" db:"logo_path"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	Status      int        `json:"status" db:"status"`
}
