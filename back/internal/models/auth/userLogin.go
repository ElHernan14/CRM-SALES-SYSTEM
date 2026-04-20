package models

import "time"

type UserLogin struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`

	CompanyID *int `json:"company_id,omitempty"` // Para clientes, puede ser nulo

	Roles []Role `json:"roles"` //Para traer roles completos

	Permissions []string `json:"permissions"` //para traer solo name permisos
}
