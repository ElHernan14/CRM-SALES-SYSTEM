package invoicedto

type GetCompanyInvoicesRequest struct {
	Page          int    `json:"page" validate:"required,min=1"`
	Limit         int    `json:"limit" validate:"required,min=1,max=100"`
	StatusInvoice string `json:"status_invoice" validate:"omitempty,max=50,status_invoice"`
	BuyerName     string `json:"buyer_name" validate:"omitempty,min=2,max=100"`
	SellerCompany string `json:"seller_company" validate:"omitempty,min=2,max=100"`
	Status        int    `json:"status" validate:"omitempty,min=0"`
	SortColumn    string `json:"sort_column" validate:"omitempty,max=50,oneof=created_at total_amount paid_amount status_invoice"`
	Order         string `json:"order" validate:"omitempty,oneof=asc desc"`
}
