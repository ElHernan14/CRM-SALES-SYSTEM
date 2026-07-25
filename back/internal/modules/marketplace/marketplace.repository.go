package marketplace

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	marketplacedto "crm-system-sales/internal/modules/marketplace/dto"
)

type MarketplaceRepository interface {
	ListSuppliers(ctx context.Context, req *marketplacedto.GetSuppliersRequest, excludedCompanyID *int) ([]marketplacedto.SupplierResponse, int, error)
	GetSupplierByID(ctx context.Context, supplierID int, excludedCompanyID *int) (*marketplacedto.SupplierResponse, error)
}

type marketplaceRepository struct {
	db *sql.DB
}

func NewMarketplaceRepository(db *sql.DB) MarketplaceRepository {
	return &marketplaceRepository{db: db}
}

func (r *marketplaceRepository) ListSuppliers(ctx context.Context, req *marketplacedto.GetSuppliersRequest, excludedCompanyID *int) ([]marketplacedto.SupplierResponse, int, error) {
	allowedSortColumns := map[string]string{
		"name":           "c.name",
		"category":       "cc.name",
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
			c.category_id,
			cc.name,
			MIN(cl.email) AS email,
			c.logo_path,
			c.cover_image_path,
			c.description,
			COUNT(p.id) AS total_products
	`

	baseQuery := `
		FROM company c
		INNER JOIN category_company cc ON cc.id = c.category_id
		LEFT JOIN client cl ON cl.company_id = c.id AND cl.status = 1 AND cl.deleted_at IS NULL
		LEFT JOIN product p ON p.company_id = c.id AND p.status = 1 AND p.deleted_at IS NULL AND p.stock > p.reserved_stock
		WHERE c.status = 1
		  AND c.deleted_at IS NULL
	`

	args := []interface{}{}
	argPos := 1

	if excludedCompanyID != nil {
		baseQuery += fmt.Sprintf(" AND c.id <> $%d", argPos)
		args = append(args, *excludedCompanyID)
		argPos++
	}

	if req.Search != "" {
		baseQuery += fmt.Sprintf(" AND c.name ILIKE $%d", argPos)
		args = append(args, "%"+req.Search+"%")
		argPos++
	}

	if req.CategoryID != nil {
		baseQuery += fmt.Sprintf(" AND c.category_id = $%d", argPos)
		args = append(args, *req.CategoryID)
		argPos++
	} else if req.Category != "" {
		baseQuery += fmt.Sprintf(" AND LOWER(cc.name) = LOWER($%d)", argPos)
		args = append(args, req.Category)
		argPos++
	}

	groupBy := `
		GROUP BY c.id, c.name, c.category_id, cc.name, c.logo_path, c.cover_image_path, c.description
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
		var email, logo, coverImage, description sql.NullString

		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.CategoryID,
			&item.Category,
			&email,
			&logo,
			&coverImage,
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
		if coverImage.Valid {
			item.CoverImage = &coverImage.String
		}
		if description.Valid {
			item.Description = &description.String
		}

		items = append(items, item)
	}

	return items, total, rows.Err()
}

func (r *marketplaceRepository) GetSupplierByID(ctx context.Context, supplierID int, excludedCompanyID *int) (*marketplacedto.SupplierResponse, error) {
	query := `
		SELECT
			c.id,
			c.name,
			c.category_id,
			cc.name,
			MIN(cl.email) AS email,
			c.logo_path,
			c.cover_image_path,
			c.description,
			COUNT(p.id) AS total_products
		FROM company c
		INNER JOIN category_company cc ON cc.id = c.category_id
		LEFT JOIN client cl ON cl.company_id = c.id AND cl.status = 1 AND cl.deleted_at IS NULL
		LEFT JOIN product p ON p.company_id = c.id AND p.status = 1 AND p.deleted_at IS NULL AND p.stock > p.reserved_stock
		WHERE c.id = $1
		  AND c.status = 1
		  AND c.deleted_at IS NULL
	`

	args := []interface{}{supplierID}
	if excludedCompanyID != nil {
		query += " AND c.id <> $2"
		args = append(args, *excludedCompanyID)
	}

	query += `
		GROUP BY c.id, c.name, c.category_id, cc.name, c.logo_path, c.cover_image_path, c.description
	`

	var item marketplacedto.SupplierResponse
	var email, logo, coverImage, description sql.NullString

	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&item.ID,
		&item.Name,
		&item.CategoryID,
		&item.Category,
		&email,
		&logo,
		&coverImage,
		&description,
		&item.TotalProducts,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if email.Valid {
		item.Email = &email.String
	}
	if logo.Valid {
		item.Logo = &logo.String
	}
	if coverImage.Valid {
		item.CoverImage = &coverImage.String
	}
	if description.Valid {
		item.Description = &description.String
	}

	return &item, nil
}
