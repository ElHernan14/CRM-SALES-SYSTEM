package dto

type EnsureCartRequest struct {
	SellerCompanyID int `json:"seller_company_id" validate:"required,gt=0"`
}

type EnsureCartResponse struct {
	InvoiceID       int    `json:"invoice_id"`
	BuyerClientID   int    `json:"buyer_client_id"`
	SellerCompanyID int    `json:"seller_company_id"`
	StatusInvoice   string `json:"status_invoice"`
}

type CartItemResponse struct {
	ID          int     `json:"id"`
	InvoiceID   int     `json:"invoice_id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Subtotal    float64 `json:"subtotal"`
}

type CartResponse struct {
	InvoiceID       int                `json:"invoice_id"`
	BuyerClientID   int                `json:"buyer_client_id"`
	SellerCompanyID int                `json:"seller_company_id"`
	SellerCompany   string             `json:"seller_company"`
	StatusInvoice   string             `json:"status_invoice"`
	Subtotal        float64            `json:"subtotal"`
	Taxes           float64            `json:"taxes"`
	TotalAmount     float64            `json:"total_amount"`
	Items           []CartItemResponse `json:"items"`
}
