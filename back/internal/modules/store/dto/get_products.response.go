package dto

import "crm-system-sales/internal/core/dto"

type StoreProductResponse struct {
	ID int `json:"id" example:"1"`

	Name string `json:"name" example:"Notebook Lenovo ThinkPad"`

	Type string `json:"type" example:"physical"`

	Description string `json:"description" example:"Business notebook"`

	Price float64 `json:"price" example:"1200"`

	AvailableStock int `json:"available_stock" example:"25"`
}

type GetStoreProductsResponse struct {
	Items []StoreProductResponse `json:"items"`

	Meta dto.Meta `json:"meta"`
}
