package marketplace

import (
	"context"
	marketplacedto "crm-system-sales/internal/modules/marketplace/dto"
	"database/sql"
	"fmt"
	"strings"
)

type MarketplaceRepository interface {
	ListSuppliers(ctx context.Context, req *marketplacedto.GetSuppliersRequest) ([]marketplacedto.SupplierResponse, int, error)
}

type marketplaceRepository struct {
	db *sql.DB
}

func NewMarketplaceRepository(db *sql.DB) MarketplaceRepository {
	return &marketplaceRepository{db: db}
}

func (r *marketplaceRepository) ListSuppliers(ctx context.Context, req *marketplacedto.GetSuppliersRequest) ([]marketplacedto.SupplierResponse, int, error) {
	allowedSortColumns := map[string]string{
		"name":           "c.name",
		"total_products": "total_products",
		"created_at":     "c.created_at",
	}

	allowedOrders := map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}

	selectQuery := `
		SELECT
			c.id,
			c.name,
			MIN(u.email) AS email,
			c.logo_path,
			c.description,
			COUNT(p.id) AS total_products
	`

	baseQuery := `
		FROM company c
		LEFT JOIN users u ON u.company_id = c.id AND u.status = 1 AND u.deleted_at IS NULL
		LEFT JOIN product p ON p.company_id = c.id AND p.status = 1 AND p.deleted_at IS NULL AND p.stock > p.reserved_stock
		WHERE c.status = 1
		  AND c.deleted_at IS NULL
	`

	args := []interface{}{}
	argPos := 1

	if req.Search != "" {
		baseQuery += fmt.Sprintf(" AND c.name ILIKE $%d", argPos)
		args = append(args, "%"+req.Search+"%")
		argPos++
	}

	groupBy := `
		GROUP BY c.id, c.name, MIN(u.email) AS email, c.logo_path, c.description
	`

	countQuery := "SELECT COUNT(*) FROM (" + selectQuery + baseQuery + groupBy + ") AS suppliers"
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	sortColumn := "c.name"
	order := "ASC"

	if v, ok := allowedSortColumns[req.SortColumn]; ok {
		sortColumn = v
	}
	if v, ok := allowedOrders[strings.ToLower(req.Order)]; ok {
		order = v
	}

	offset := (req.Page - 1) * req.Limit
	query := selectQuery + baseQuery + groupBy
	query += fmt.Sprintf(" ORDER BY %s %s", sortColumn, order)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)

	args = append(args, req.Limit, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]marketplacedto.SupplierResponse, 0)
	for rows.Next() {
		var item marketplacedto.SupplierResponse
		var email, logo, description sql.NullString

		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&email,
			&logo,
			&description,
			&item.TotalProducts,
		); err != nil {
			return nil, 0, err
		}

		if email.Valid {
			item.Email = &email.String
		}
		if logo.Valid {
			item.Logo = &logo.String
		}
		if description.Valid {
			item.Description = &description.String
		}

		items = append(items, item)
	}

	return items, total, rows.Err()
}
