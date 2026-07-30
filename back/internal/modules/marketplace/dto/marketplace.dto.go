package marketplacedto

import "crm-system-sales/internal/core/dto"

type GetSuppliersRequest struct {
	Search      string `json:"search" validate:"omitempty,max=100"`
	CategoryID  *int   `json:"category_id" validate:"omitempty,gt=0"`
	Category    string `json:"category" validate:"omitempty,min=2,max=80"`
	Page        int    `json:"page" validate:"required,min=1"`
	Limit       int    `json:"limit" validate:"required,min=1,max=100"`
	SortColumn  string `json:"sort_column" validate:"omitempty,max=50,oneof=name category total_products created_at"`
	Order       string `json:"order" validate:"omitempty,oneof=asc desc"`
	IncludeSelf bool   `json:"include_self" validate:"-"`
}

type SupplierResponse struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	CategoryID    int     `json:"category_id"`
	Category      string  `json:"category"`
	Email         *string `json:"email,omitempty"`
	Logo          *string `json:"logo,omitempty"`
	CoverImage    *string `json:"cover_image,omitempty"`
	Description   *string `json:"description,omitempty"`
	TotalProducts int     `json:"total_products"`
}

type GetSuppliersResponse struct {
	Items []SupplierResponse `json:"items"`
	Meta  dto.Meta           `json:"meta"`
}

type EnsureMarketplaceCartRequest struct {
	SellerCompanyID int `json:"seller_company_id" validate:"required,gt=0"`
}

type EnsureMarketplaceCartResponse struct {
	InvoiceID       int    `json:"invoice_id"`
	BuyerClientID   int    `json:"buyer_client_id"`
	SellerCompanyID int    `json:"seller_company_id"`
	StatusInvoice   string `json:"status_invoice"`
	Source          string `json:"source"`
}

type MarketplaceCartItemResponse struct {
	ID               int     `json:"id"`
	InvoiceID        int     `json:"invoice_id"`
	ProductID        int     `json:"product_id"`
	ProductName      string  `json:"product_name"`
	ProductImagePath *string `json:"product_image_path,omitempty"`
	Quantity         int     `json:"quantity"`
	Price            float64 `json:"price"`
	Subtotal         float64 `json:"subtotal"`
}

type MarketplaceCartResponse struct {
	InvoiceID       int                           `json:"invoice_id"`
	BuyerClientID   int                           `json:"buyer_client_id"`
	SellerCompanyID int                           `json:"seller_company_id"`
	SellerCompany   string                        `json:"seller_company"`
	StatusInvoice   string                        `json:"status_invoice"`
	Source          string                        `json:"source"`
	Subtotal        float64                       `json:"subtotal"`
	Taxes           float64                       `json:"taxes"`
	TotalAmount     float64                       `json:"total_amount"`
	Items           []MarketplaceCartItemResponse `json:"items"`
}

type MarketplaceCheckoutRequest struct {
	InvoiceID       *int `json:"invoice_id,omitempty" validate:"omitempty,gt=0"`
	SellerCompanyID *int `json:"seller_company_id,omitempty" validate:"omitempty,gt=0"`
}

type MarketplaceCheckoutResponse struct {
	InvoiceID int    `json:"invoice_id"`
	Status    string `json:"status"`
	Source    string `json:"source"`
}
