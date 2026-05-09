package model

type Invoice struct {
	ID              int    `db:"id"`
	BuyerClientID   int    `db:"buyer_client_id"`
	SellerCompanyID int    `db:"seller_company_id"`
	CreatedByUserID int    `db:"created_by_user_id"`
	StatusInvoice   string `db:"status_invoice"`
}
