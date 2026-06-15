package invoiceitemdto

type CreateInvoiceItemRequest struct {
	ProductID int `json:"product_id" validate:"required,gt=0" example:"12"`
	Quantity  int `json:"quantity" validate:"required,gt=0,lte=999" example:"3"`
}
