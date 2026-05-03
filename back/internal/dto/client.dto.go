package dto

type CreateClientRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"strong_password,required,min=6"`
	Phone     string `json:"phone" validate:"omitempty,max=50"`
	CompanyID *int   `json:"company_id,omitempty"`
}

type ClientResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CompanyID *int   `json:"company_id,omitempty"`
}

type GetClientsRequest struct {
	CompanyID *int   `validate:"omitempty,gt=0"`
	Search    string `validate:"omitempty,min=2,max=100"`
	Email     string `validate:"omitempty,email"`
	Page      int    `validate:"gte=1"`
	Limit     int    `validate:"gte=1,lte=100"`
}

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type GetClientsResponse struct {
	Clients []ClientResponse `json:"clients"`
	Meta    Meta             `json:"meta"`
}
