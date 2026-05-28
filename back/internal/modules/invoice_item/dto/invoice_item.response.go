package invoiceitemdto

type InvoiceItemResponse struct {
	ID int `json:"id"`

	InvoiceID *int `json:"invoice_id,omitempty"`

	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`

	Quantity int `json:"quantity"`

	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`
}
