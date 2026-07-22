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
	ID               int     `json:"id"`
	InvoiceID        int     `json:"invoice_id"`
	ProductID        int     `json:"product_id"`
	ProductName      string  `json:"product_name"`
	ProductImagePath *string `json:"product_image_path,omitempty"`
	Quantity         int     `json:"quantity"`
	Price            float64 `json:"price"`
	Subtotal         float64 `json:"subtotal"`
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

type CartsSummaryResponse struct {
	SellerCount int     `json:"seller_count"`
	ItemCount   int     `json:"item_count"`
	Subtotal    float64 `json:"subtotal"`
	Taxes       float64 `json:"taxes"`
	TotalAmount float64 `json:"total_amount"`
}

type GetCartsResponse struct {
	Carts   []CartResponse       `json:"carts"`
	Summary CartsSummaryResponse `json:"summary"`
}

type CheckoutAllRequest struct {
	InvoiceIDs []int `json:"invoice_ids,omitempty" validate:"omitempty,min=1"`
}

type CheckoutAllOrderResponse struct {
	InvoiceID       int    `json:"invoice_id"`
	SellerCompanyID int    `json:"seller_company_id"`
	SellerCompany   string `json:"seller_company"`
	Status          string `json:"status"`
}

type CheckoutAllResponse struct {
	Orders      []CheckoutAllOrderResponse `json:"orders"`
	Count       int                        `json:"count"`
	TotalAmount float64                    `json:"total_amount"`
}
