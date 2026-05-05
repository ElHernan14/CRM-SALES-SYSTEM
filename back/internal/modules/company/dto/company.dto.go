package companydto

import "crm-system-sales/internal/core/dto"

type CreateCompanyRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,strong_password"`
}

type CompanyResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	CreatedAt *string `json:"created_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Status    *int    `json:"status,omitempty"`
}

type GetCompaniesRequest struct {
	Search string `validate:"omitempty,min=2,max=100"`
	Page   int    `validate:"gte=1"`
	Limit  int    `validate:"gte=1,lte=100"`
}

type CompanyListItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetCompaniesResponse struct {
	Data []CompanyListItem `json:"data"`
	Meta dto.Meta          `json:"meta"`
}
