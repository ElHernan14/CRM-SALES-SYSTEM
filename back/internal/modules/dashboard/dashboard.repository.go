package dashboard

import (
	"context"
	"database/sql"

	dashboarddto "crm-system-sales/internal/modules/dashboard/dto"
)

type DashboardRepository interface {
	GetSalesSummary(ctx context.Context, companyID int) (dashboarddto.DashboardSalesSummaryResponse, error)
	GetPurchasesSummary(ctx context.Context, companyID int) (dashboarddto.DashboardPurchasesSummaryResponse, error)
	GetInventorySummary(ctx context.Context, companyID int) (dashboarddto.DashboardInventorySummaryResponse, error)
	CountPartiallyPaidPurchases(ctx context.Context, companyID int) (int, error)
	GetRecentSales(ctx context.Context, companyID int, limit int) ([]dashboarddto.DashboardRecentInvoiceResponse, error)
	GetRecentPurchases(ctx context.Context, companyID int, limit int) ([]dashboarddto.DashboardRecentInvoiceResponse, error)
	GetLowStockProducts(ctx context.Context, companyID int, limit int) ([]dashboarddto.DashboardLowStockProductResponse, error)
}

type dashboardRepository struct {
	db *sql.DB
}

func NewDashboardRepository(db *sql.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

func (r *dashboardRepository) GetSalesSummary(ctx context.Context, companyID int) (dashboarddto.DashboardSalesSummaryResponse, error) {
	query := `
		SELECT
			COUNT(*)::int AS total_invoices,
			COALESCE(SUM(CASE WHEN i.status_invoice = 'draft' THEN 1 ELSE 0 END), 0)::int AS draft,
			COALESCE(SUM(CASE WHEN i.status_invoice = 'pending' THEN 1 ELSE 0 END), 0)::int AS pending,
			COALESCE(SUM(CASE WHEN i.status_invoice = 'paid' THEN 1 ELSE 0 END), 0)::int AS paid,
			COALESCE(SUM(CASE WHEN i.status_invoice = 'cancelled' THEN 1 ELSE 0 END), 0)::int AS cancelled,
			COALESCE(SUM(i.total_amount), 0) AS total_amount,
			COALESCE(SUM(i.paid_amount), 0) AS paid_amount,
			ROUND(COALESCE(SUM(CASE WHEN i.status_invoice = 'pending' THEN GREATEST(i.total_amount - i.paid_amount, 0) ELSE 0 END), 0), 2) AS outstanding_amount
		FROM invoice i
		WHERE i.seller_company_id = $1
		  AND i.status = 1
		  AND i.deleted_at IS NULL
	`

	var res dashboarddto.DashboardSalesSummaryResponse
	err := r.db.QueryRowContext(ctx, query, companyID).Scan(
		&res.TotalInvoices,
		&res.Draft,
		&res.Pending,
		&res.Paid,
		&res.Cancelled,
		&res.TotalAmount,
		&res.PaidAmount,
		&res.OutstandingAmount,
	)
	return res, err
}

func (r *dashboardRepository) GetPurchasesSummary(ctx context.Context, companyID int) (dashboarddto.DashboardPurchasesSummaryResponse, error) {
	query := `
		SELECT
			COUNT(*)::int AS total_invoices,
			COALESCE(SUM(CASE WHEN i.status_invoice = 'draft' THEN 1 ELSE 0 END), 0)::int AS draft,
			COALESCE(SUM(CASE WHEN i.status_invoice = 'pending' THEN 1 ELSE 0 END), 0)::int AS pending,
			COALESCE(SUM(CASE WHEN i.status_invoice = 'paid' THEN 1 ELSE 0 END), 0)::int AS paid,
			COALESCE(SUM(CASE WHEN i.status_invoice = 'cancelled' THEN 1 ELSE 0 END), 0)::int AS cancelled,
			COALESCE(SUM(i.total_amount), 0) AS total_amount,
			COALESCE(SUM(i.paid_amount), 0) AS paid_amount,
			ROUND(COALESCE(SUM(CASE WHEN i.status_invoice = 'pending' THEN GREATEST(i.total_amount - i.paid_amount, 0) ELSE 0 END), 0), 2) AS outstanding_amount
		FROM invoice i
		INNER JOIN client buyer ON buyer.id = i.buyer_client_id
		WHERE buyer.company_id = $1
		  AND COALESCE(i.source, 'erp') = 'erp'
		  AND i.status = 1
		  AND i.deleted_at IS NULL
	`

	var res dashboarddto.DashboardPurchasesSummaryResponse
	err := r.db.QueryRowContext(ctx, query, companyID).Scan(
		&res.TotalInvoices,
		&res.Draft,
		&res.Pending,
		&res.Paid,
		&res.Cancelled,
		&res.TotalAmount,
		&res.PaidAmount,
		&res.OutstandingAmount,
	)
	return res, err
}

func (r *dashboardRepository) GetInventorySummary(ctx context.Context, companyID int) (dashboarddto.DashboardInventorySummaryResponse, error) {
	query := `
		SELECT
			COALESCE(SUM(CASE WHEN p.status = 1 AND p.deleted_at IS NULL THEN 1 ELSE 0 END), 0)::int AS active_products,
			COALESCE(SUM(CASE WHEN p.status = 0 OR p.deleted_at IS NOT NULL THEN 1 ELSE 0 END), 0)::int AS inactive_products,
			COALESCE(SUM(CASE WHEN p.status = 1 AND p.deleted_at IS NULL THEN GREATEST(p.stock - p.reserved_stock, 0) ELSE 0 END), 0)::int AS available_units,
			COALESCE(SUM(CASE WHEN p.status = 1 AND p.deleted_at IS NULL THEN p.reserved_stock ELSE 0 END), 0)::int AS reserved_units,
			COALESCE(SUM(CASE WHEN p.status = 1 AND p.deleted_at IS NULL AND (p.stock - p.reserved_stock) > 0 AND (p.stock - p.reserved_stock) <= 10 THEN 1 ELSE 0 END), 0)::int AS low_stock_products,
			COALESCE(SUM(CASE WHEN p.status = 1 AND p.deleted_at IS NULL AND (p.stock - p.reserved_stock) <= 0 THEN 1 ELSE 0 END), 0)::int AS out_of_stock_products
		FROM product p
		WHERE p.company_id = $1
	`

	var res dashboarddto.DashboardInventorySummaryResponse
	err := r.db.QueryRowContext(ctx, query, companyID).Scan(
		&res.ActiveProducts,
		&res.InactiveProducts,
		&res.AvailableUnits,
		&res.ReservedUnits,
		&res.LowStockProducts,
		&res.OutOfStockProducts,
	)
	return res, err
}

func (r *dashboardRepository) CountPartiallyPaidPurchases(ctx context.Context, companyID int) (int, error) {
	query := `
		SELECT COUNT(*)::int
		FROM invoice i
		INNER JOIN client buyer ON buyer.id = i.buyer_client_id
		WHERE buyer.company_id = $1
		  AND COALESCE(i.source, 'erp') = 'erp'
		  AND i.status = 1
		  AND i.deleted_at IS NULL
		  AND i.status_invoice = 'pending'
		  AND i.paid_amount > 0
		  AND i.paid_amount < i.total_amount
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, companyID).Scan(&count)
	return count, err
}

func (r *dashboardRepository) GetRecentSales(ctx context.Context, companyID int, limit int) ([]dashboarddto.DashboardRecentInvoiceResponse, error) {
	query := `
		SELECT
			i.id,
			COALESCE(NULLIF(buyer_company.name, ''), NULLIF(TRIM(COALESCE(buyer.first_name, '') || ' ' || COALESCE(buyer.last_name, '')), ''), 'Unknown') AS counterparty,
			i.status_invoice,
			i.total_amount,
			i.paid_amount,
			ROUND(GREATEST(i.total_amount - i.paid_amount, 0), 2) AS remaining_amount,
			i.created_at
		FROM invoice i
		LEFT JOIN client buyer ON buyer.id = i.buyer_client_id
		LEFT JOIN company buyer_company ON buyer_company.id = buyer.company_id
		WHERE i.seller_company_id = $1
		  AND i.status = 1
		  AND i.deleted_at IS NULL
		ORDER BY i.created_at DESC
		LIMIT $2
	`

	return r.queryRecentInvoices(ctx, query, companyID, limit)
}

func (r *dashboardRepository) GetRecentPurchases(ctx context.Context, companyID int, limit int) ([]dashboarddto.DashboardRecentInvoiceResponse, error) {
	query := `
		SELECT
			i.id,
			COALESCE(NULLIF(seller.name, ''), 'Unknown') AS counterparty,
			i.status_invoice,
			i.total_amount,
			i.paid_amount,
			ROUND(GREATEST(i.total_amount - i.paid_amount, 0), 2) AS remaining_amount,
			i.created_at
		FROM invoice i
		INNER JOIN client buyer ON buyer.id = i.buyer_client_id
		LEFT JOIN company seller ON seller.id = i.seller_company_id
		WHERE buyer.company_id = $1
		  AND COALESCE(i.source, 'erp') = 'erp'
		  AND i.status = 1
		  AND i.deleted_at IS NULL
		ORDER BY i.created_at DESC
		LIMIT $2
	`

	return r.queryRecentInvoices(ctx, query, companyID, limit)
}

func (r *dashboardRepository) queryRecentInvoices(ctx context.Context, query string, companyID int, limit int) ([]dashboarddto.DashboardRecentInvoiceResponse, error) {
	rows, err := r.db.QueryContext(ctx, query, companyID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoices := make([]dashboarddto.DashboardRecentInvoiceResponse, 0)
	for rows.Next() {
		var inv dashboarddto.DashboardRecentInvoiceResponse
		if err := rows.Scan(
			&inv.ID,
			&inv.Counterparty,
			&inv.StatusInvoice,
			&inv.TotalAmount,
			&inv.PaidAmount,
			&inv.RemainingAmount,
			&inv.CreatedAt,
		); err != nil {
			return nil, err
		}
		invoices = append(invoices, inv)
	}

	return invoices, rows.Err()
}

func (r *dashboardRepository) GetLowStockProducts(ctx context.Context, companyID int, limit int) ([]dashboarddto.DashboardLowStockProductResponse, error) {
	query := `
		SELECT
			p.id,
			p.name,
			cp.name AS category,
			pt.name AS type,
			p.stock,
			p.reserved_stock,
			p.stock - p.reserved_stock AS available_stock,
			p.image_path
		FROM product p
		INNER JOIN category_product cp ON cp.id = p.category_id
		INNER JOIN product_type pt ON pt.id = p.type_id
		WHERE p.company_id = $1
		  AND p.status = 1
		  AND p.deleted_at IS NULL
		  AND (p.stock - p.reserved_stock) <= 10
		ORDER BY (p.stock - p.reserved_stock) ASC, p.name ASC
		LIMIT $2
	`

	rows, err := r.db.QueryContext(ctx, query, companyID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]dashboarddto.DashboardLowStockProductResponse, 0)
	for rows.Next() {
		var p dashboarddto.DashboardLowStockProductResponse
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Category,
			&p.Type,
			&p.Stock,
			&p.ReservedStock,
			&p.AvailableStock,
			&p.ImagePath,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, rows.Err()
}
