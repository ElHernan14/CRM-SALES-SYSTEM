package clientdto

import "crm-system-sales/internal/core/dto"

type CompanyCustomerResponse struct {
	ID int `json:"id"`

	Name string `json:"name"`

	Email string `json:"email"`

	CompanyID *int `json:"company_id"`

	TotalInvoices int `json:"total_invoices"`

	TotalPurchased float64 `json:"total_purchased"`
}

type GetCompanyCustomersResponse struct {
	Items []CompanyCustomerResponse `json:"items"`

	Meta dto.Meta `json:"meta"`
}
