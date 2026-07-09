package dto

import "crm-system-sales/internal/core/dto"

type StoreProductResponse struct {
	ID int `json:"id" example:"1"`

	CompanyID int `json:"company_id" example:"2"`

	CompanyName string `json:"company_name,omitempty" example:"TechNova Inc."`

	Name string `json:"name" example:"Notebook Lenovo ThinkPad"`

	Type string `json:"type" example:"physical"`

	Description string `json:"description" example:"Business notebook"`

	Price float64 `json:"price" example:"1200"`

	AvailableStock int `json:"available_stock" example:"25"`

	ImagePath *string `json:"image_path,omitempty" example:"/uploads/product/images/product_1.png"`
}

type GetStoreProductsResponse struct {
	Items []StoreProductResponse `json:"items"`

	Meta dto.Meta `json:"meta"`
}
