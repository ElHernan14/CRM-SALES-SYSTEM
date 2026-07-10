package productdto

import "crm-system-sales/internal/core/dto"

type CompanyProductResponse struct {
	ID int `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Type string `json:"type"`

	CategoryID int `json:"category_id"`

	Category string `json:"category"`

	Price float64 `json:"price"`

	Stock int `json:"stock"`

	ReservedStock int `json:"reserved_stock"`

	Status int `json:"status"`

	ImagePath *string `json:"image_path,omitempty"`
}

type GetCompanyProductsResponse struct {
	Items []CompanyProductResponse `json:"items"`

	Meta dto.Meta `json:"meta"`
}
