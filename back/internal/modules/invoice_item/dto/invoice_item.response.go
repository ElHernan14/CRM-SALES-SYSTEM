package invoiceitemdto

type InvoiceItemResponse struct {
	ID int `json:"id" example:"101"`

	InvoiceID *int `json:"invoice_id,omitempty" example:"15"`

	ProductID   int    `json:"product_id" example:"12"`
	ProductName string `json:"product_name" example:"Notebook Lenovo ThinkPad"`

	Quantity int `json:"quantity" example:"3"`

	Price    float64 `json:"price" example:"1200"`
	Subtotal float64 `json:"subtotal" example:"3600"`
}
