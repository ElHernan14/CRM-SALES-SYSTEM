package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`

	Roles []Role `json:"roles"` //Para traer roles completos

	Permissions []string `json:"permissions"` //para traer solo name permisos
}
