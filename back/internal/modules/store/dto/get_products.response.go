package dto

import "crm-system-sales/internal/core/dto"

type StoreProductResponse struct {
	ID int `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Price float64 `json:"price"`

	AvailableStock int `json:"available_stock"`
}

type GetStoreProductsResponse struct {
	Items []StoreProductResponse `json:"items"`

	Meta dto.Meta `json:"meta"`
}
