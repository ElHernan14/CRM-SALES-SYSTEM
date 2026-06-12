package invoicedto

import (
	"crm-system-sales/internal/core/dto"
	"time"
)

type CompanyInvoiceResponse struct {
	ID int `json:"id"`

	BuyerClientID int `json:"buyer_client_id"`

	SellerCompanyID int `json:"seller_company_id"`

	StatusInvoice string `json:"status_invoice"`

	TotalAmount float64 `json:"total_amount"`

	PaidAmount float64 `json:"paid_amount"`

	CreatedAt time.Time `json:"created_at"`
}

type GetCompanyInvoicesResponse struct {
	Items []CompanyInvoiceResponse `json:"items"`

	Meta dto.Meta `json:"meta"`
}
