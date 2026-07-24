package dashboarddto

import "time"

type DashboardSalesSummaryResponse struct {
	TotalInvoices     int     `json:"total_invoices"`
	Draft             int     `json:"draft"`
	Pending           int     `json:"pending"`
	Paid              int     `json:"paid"`
	Cancelled         int     `json:"cancelled"`
	TotalAmount       float64 `json:"total_amount"`
	PaidAmount        float64 `json:"paid_amount"`
	OutstandingAmount float64 `json:"outstanding_amount"`
}

type DashboardPurchasesSummaryResponse struct {
	TotalInvoices     int     `json:"total_invoices"`
	Draft             int     `json:"draft"`
	Pending           int     `json:"pending"`
	Paid              int     `json:"paid"`
	Cancelled         int     `json:"cancelled"`
	TotalAmount       float64 `json:"total_amount"`
	PaidAmount        float64 `json:"paid_amount"`
	OutstandingAmount float64 `json:"outstanding_amount"`
}

type DashboardInventorySummaryResponse struct {
	ActiveProducts     int `json:"active_products"`
	InactiveProducts   int `json:"inactive_products"`
	AvailableUnits     int `json:"available_units"`
	ReservedUnits      int `json:"reserved_units"`
	LowStockProducts   int `json:"low_stock_products"`
	OutOfStockProducts int `json:"out_of_stock_products"`
}

type DashboardRecentInvoiceResponse struct {
	ID              int       `json:"id"`
	Counterparty    string    `json:"counterparty"`
	StatusInvoice   string    `json:"status_invoice"`
	TotalAmount     float64   `json:"total_amount"`
	PaidAmount      float64   `json:"paid_amount"`
	RemainingAmount float64   `json:"remaining_amount"`
	CreatedAt       time.Time `json:"created_at"`
}

type DashboardLowStockProductResponse struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Type           string  `json:"type"`
	Stock          int     `json:"stock"`
	ReservedStock  int     `json:"reserved_stock"`
	AvailableStock int     `json:"available_stock"`
	ImagePath      *string `json:"image_path,omitempty"`
}

type DashboardAttentionResponse struct {
	PendingSales           int `json:"pending_sales"`
	DraftSales             int `json:"draft_sales"`
	PendingPurchases       int `json:"pending_purchases"`
	PartiallyPaidPurchases int `json:"partially_paid_purchases"`
	LowStockProducts       int `json:"low_stock_products"`
	OutOfStockProducts     int `json:"out_of_stock_products"`
}

type DashboardOverviewResponse struct {
	Sales     DashboardSalesSummaryResponse     `json:"sales"`
	Purchases DashboardPurchasesSummaryResponse `json:"purchases"`
	Inventory DashboardInventorySummaryResponse `json:"inventory"`
	Attention DashboardAttentionResponse        `json:"attention"`

	RecentSales      []DashboardRecentInvoiceResponse   `json:"recent_sales"`
	RecentPurchases  []DashboardRecentInvoiceResponse   `json:"recent_purchases"`
	LowStockProducts []DashboardLowStockProductResponse `json:"low_stock_products"`
}
