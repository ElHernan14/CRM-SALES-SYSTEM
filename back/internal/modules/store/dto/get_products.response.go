package dto

import "crm-system-sales/internal/core/dto"

type StoreProductResponse struct {
	ID int `json:"id" example:"1"`

	CompanyID int `json:"company_id" example:"2"`

	CompanyName string `json:"company_name,omitempty" example:"TechNova Inc."`

	Name string `json:"name" example:"Notebook Lenovo ThinkPad"`

	Kind string `json:"kind" example:"product"`

	CategoryID int `json:"category_id" example:"1"`

	Category string `json:"category" example:"Technology"`

	TypeID int `json:"type_id" example:"3"`

	Type string `json:"type" example:"Notebook"`

	Description string `json:"description" example:"Business notebook"`

	Price float64 `json:"price" example:"1200"`

	AvailableStock int `json:"available_stock" example:"25"`

	ImagePath *string `json:"image_path,omitempty" example:"/uploads/product/images/product_1.png"`
}

type GetStoreProductsResponse struct {
	Items []StoreProductResponse `json:"items"`

	Meta dto.Meta `json:"meta"`
}

type StoreProductDetailResponse struct {
	ID int `json:"id" example:"1"`

	CompanyID int `json:"company_id" example:"2"`

	CompanyName string `json:"company_name" example:"TechNova"`

	Name string `json:"name" example:"Notebook Lenovo"`

	Description string `json:"description" example:"Business notebook"`

	Kind string `json:"kind" example:"product"`

	CategoryID int `json:"category_id" example:"1"`

	Category string `json:"category" example:"Technology"`

	TypeID int `json:"type_id" example:"3"`

	Type string `json:"type" example:"Notebook"`

	Price float64 `json:"price" example:"1200"`

	AvailableStock int `json:"available_stock" example:"25"`

	ImagePath *string `json:"image_path,omitempty" example:"/uploads/product/images/product_1.png"`

	IsAvailable bool `json:"is_available" example:"true"`

	UnavailableReason string `json:"unavailable_reason,omitempty" example:"product_unavailable"`
}
