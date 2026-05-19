package invoiceitemdto

type UpdateInvoiceItemRequest struct {
	Quantity int `json:"quantity" validate:"required,gt=0,lte=9999"`
}
