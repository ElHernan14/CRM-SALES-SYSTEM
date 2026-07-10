package models

import "time"

type Company struct {
	ID             int        `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	CategoryID     int        `json:"category_id" db:"category_id"`
	Category       string     `json:"category" db:"category"`
	Description    *string    `json:"description,omitempty" db:"description"`
	LogoPath       *string    `json:"logo_path,omitempty" db:"logo_path"`
	CoverImagePath *string    `json:"cover_image_path,omitempty" db:"cover_image_path"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	Status         int        `json:"status" db:"status"`
}
