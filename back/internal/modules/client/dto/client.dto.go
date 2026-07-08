package clientdto

import "crm-system-sales/internal/core/dto"

type CreateClientRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"strong_password,required,min=6"`
	Phone     string `json:"phone" validate:"omitempty,max=50"`
	CompanyID *int   `json:"company_id,omitempty"`
}

type ClientResponse struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	CompanyID *int    `json:"company_id,omitempty"`
	Phone     string  `json:"phone"`
	Status    int     `json:"status,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
}

type GetClientsRequest struct {
	CompanyID *int   `validate:"omitempty,gt=0"`
	Search    string `validate:"omitempty,min=2,max=100"`
	Email     string `validate:"omitempty,email"`
	Page      int    `validate:"gte=1"`
	Limit     int    `validate:"gte=1,lte=100"`
}

type GetClientsResponse struct {
	Clients []ClientResponse `json:"clients"`
	Meta    dto.Meta         `json:"meta"`
}

type UpdateClientRequest struct {
	FirstName *string `json:"first_name" validate:"omitempty,min=2,max=100"`
	LastName  *string `json:"last_name" validate:"omitempty,min=2,max=100"`
	Email     *string `json:"email" validate:"omitempty,email"`
	Phone     *string `json:"phone" validate:"omitempty,min=6,max=20"`
}
