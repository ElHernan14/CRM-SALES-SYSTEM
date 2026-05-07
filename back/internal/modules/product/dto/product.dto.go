package productdto

import (
	meta "crm-system-sales/internal/core/dto"
)

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Description string  `json:"description" validate:"omitempty,max=500"`
	Type        string  `json:"type" validate:"required,product_type"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"gte=0"`
	CompanyID   *int    `json:"company_id,omitempty"`
}

type ProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type GetProductsRequest struct {
	Search    string  `validate:"omitempty,min=2,max=100"`
	Type      string  `validate:"omitempty,product_type"`
	MinPrice  float64 `validate:"omitempty,gte=0"`
	MaxPrice  float64 `validate:"omitempty,gte=0"`
	CompanyID *int    `validate:"omitempty,gt=0"`
	Page      int     `validate:"gte=1"`
	Limit     int     `validate:"gte=1,lte=100"`
}

type ProductListItem struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      int     `json:"status"`
	CompanyID   int     `json:"company_id"`
}

type GetProductsResponse struct {
	Data []ProductListItem `json:"data"`
	Meta meta.Meta         `json:"meta"`
}

type ProductDetailResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      int     `json:"status"`
	CompanyID   int     `json:"company_id"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name" validate:"omitempty,min=2,max=100"`
	Description *string  `json:"description" validate:"omitempty,max=500"`
	Type        *string  `json:"type" validate:"omitempty,product_type"`
	Price       *float64 `json:"price" validate:"omitempty,gte=0"`
	Stock       *int     `json:"stock" validate:"omitempty,gte=0"`
	Status      *int     `json:"status" validate:"omitempty,oneof=0 1"`
}
