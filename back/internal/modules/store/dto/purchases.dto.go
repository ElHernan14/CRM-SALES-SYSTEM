package dto

import (
	metadto "crm-system-sales/internal/core/dto"
	"time"
)

type GetStorePurchasesRequest struct {
	Page  int `json:"page" validate:"required,min=1"`
	Limit int `json:"limit" validate:"required,min=1,max=100"`

	Search          string `json:"search" validate:"omitempty,max=100"`
	StatusInvoice   string `json:"status_invoice" validate:"omitempty,status_invoice"`
	SellerCompanyID *int   `json:"seller_company_id" validate:"omitempty,gt=0"`

	SortColumn string `json:"sort_column" validate:"omitempty,oneof=created_at total_amount paid_amount status_invoice"`
	Order      string `json:"order" validate:"omitempty,oneof=asc desc"`
}

type StorePurchaseResponse struct {
	ID int `json:"id"`

	SellerCompanyID int    `json:"seller_company_id"`
	SellerCompany   string `json:"seller_company"`

	StatusInvoice string `json:"status_invoice"`

	Subtotal        float64 `json:"subtotal"`
	Taxes           float64 `json:"taxes"`
	TotalAmount     float64 `json:"total_amount"`
	PaidAmount      float64 `json:"paid_amount"`
	RemainingAmount float64 `json:"remaining_amount"`

	ItemCount int `json:"item_count"`

	CreatedAt time.Time `json:"created_at"`
}

type GetStorePurchasesResponse struct {
	Items []StorePurchaseResponse `json:"items"`
	Meta  metadto.Meta            `json:"meta"`
}
