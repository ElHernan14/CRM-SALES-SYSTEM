package dto

import coreDto "crm-system-sales/internal/core/dto"

type GetSuppliersRequest struct {
	Search     string `json:"search" validate:"omitempty,min=2,max=100"`
	Page       int    `json:"page" validate:"required,min=1"`
	Limit      int    `json:"limit" validate:"required,min=1,max=100"`
	SortColumn string `json:"sort_column" validate:"omitempty,max=50,oneof=name total_products created_at"`
	Order      string `json:"order" validate:"omitempty,oneof=asc desc"`
}

type SupplierResponse struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Email         *string `json:"email,omitempty"`
	Logo          *string `json:"logo,omitempty"`
	Description   *string `json:"description,omitempty"`
	TotalProducts int     `json:"total_products"`
}

type GetSuppliersResponse struct {
	Items []SupplierResponse `json:"items"`
	Meta  coreDto.Meta       `json:"meta"`
}
