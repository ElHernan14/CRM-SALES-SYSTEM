package invoicedto

type CreateInvoiceRequest struct {
	BuyerClientID   int `json:"buyer_client_id" validate:"required,gt=0" example:"5"`
	SellerCompanyID int `json:"seller_company_id" validate:"required,gt=0" example:"2"`
}
