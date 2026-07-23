package companydto

import "crm-system-sales/internal/core/dto"

type CreateCompanyRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	CategoryID  int     `json:"category_id" validate:"required,gt=0"`
	Email       string  `json:"email" validate:"required,email"`
	Password    string  `json:"password" validate:"required,strong_password"`
	Description *string `json:"description" validate:"omitempty,max=500"`
}

type UpdateMyCompanyRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=2,max=100"`
	CategoryID  *int    `json:"category_id" validate:"omitempty,gt=0"`
	Description *string `json:"description" validate:"omitempty,max=500"`
}

type CompanyResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CategoryID  int     `json:"category_id"`
	Category    string  `json:"category"`
	Description *string `json:"description,omitempty"`
	Logo        *string `json:"logo,omitempty"`
	CoverImage  *string `json:"cover_image,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
	Status      *int    `json:"status,omitempty"`
}

type UploadCompanyLogoResponse struct {
	Logo string `json:"logo"`
}

type UploadCompanyCoverImageResponse struct {
	CoverImage string `json:"cover_image"`
}

type GetCompaniesRequest struct {
	Search     string `validate:"omitempty,min=2,max=100"`
	CategoryID *int   `validate:"omitempty,gt=0"`
	Category   string `validate:"omitempty,min=2,max=80"`
	Page       int    `validate:"gte=1"`
	Limit      int    `validate:"gte=1,lte=100"`
}

type CompanyListItem struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	CategoryID int     `json:"category_id"`
	Category   string  `json:"category"`
	Logo       *string `json:"logo,omitempty"`
	CoverImage *string `json:"cover_image,omitempty"`
}

type GetCompaniesResponse struct {
	Data []CompanyListItem `json:"data"`
	Meta dto.Meta          `json:"meta"`
}
