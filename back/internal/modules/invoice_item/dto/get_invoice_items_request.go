package invoiceitemdto

type GetInvoiceItemsRequest struct {
	Page  int `validate:"gte=1"`
	Limit int `validate:"gte=1,lte=100"`
}
