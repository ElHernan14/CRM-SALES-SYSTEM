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
	Source        string `json:"source"`

	TotalAmount float64 `json:"total_amount"`

	PaidAmount float64 `json:"paid_amount"`

	CreatedAt time.Time `json:"created_at"`

	BuyerName     string `json:"buyer_name"`
	SellerCompany string `json:"seller_company"`
}

type GetCompanyInvoicesResponse struct {
	Items []CompanyInvoiceResponse `json:"items"`

	Meta dto.Meta `json:"meta"`
}
