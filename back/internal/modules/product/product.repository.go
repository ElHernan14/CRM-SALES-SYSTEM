package product

import (
	"context"
	models "crm-system-sales/internal/models/product"
	"database/sql"
	"fmt"
)

type ProductRepository interface {
	Create(ctx context.Context, p *models.Product) error
	GetAll(ctx context.Context, search string, productType string, minPrice float64, maxPrice float64, companyID *int, limit int, offset int) ([]models.Product, int, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, p *models.Product) error {

	query := `
		INSERT INTO product (company_id, name, description, type, price, stock, status)
		VALUES ($1, $2, $3, $4, $5, $6, 1)
		RETURNING id
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		p.CompanyID,
		p.Name,
		p.Description,
		p.Type,
		p.Price,
		p.Stock,
	).Scan(&p.ID)
}

func (r *productRepository) GetAll(
	ctx context.Context,
	search string,
	productType string,
	minPrice float64,
	maxPrice float64,
	companyID *int,
	limit int,
	offset int,
) ([]models.Product, int, error) {

	baseQuery := `
		FROM product
		WHERE deleted_at IS NULL
		  AND status = 1
	`

	var args []interface{}
	i := 1

	if companyID != nil {
		baseQuery += fmt.Sprintf(" AND company_id = $%d", i)
		args = append(args, *companyID)
		i++
	}

	if search != "" {
		baseQuery += fmt.Sprintf(`
			AND LOWER(name) LIKE LOWER($%d)
		`, i)
		args = append(args, "%"+search+"%")
		i++
	}

	if productType != "" {
		baseQuery += fmt.Sprintf(" AND type = $%d", i)
		args = append(args, productType)
		i++
	}

	if minPrice > 0 {
		baseQuery += fmt.Sprintf(" AND price >= $%d", i)
		args = append(args, minPrice)
		i++
	}

	if maxPrice > 0 {
		baseQuery += fmt.Sprintf(" AND price <= $%d", i)
		args = append(args, maxPrice)
		i++
	}

	countQuery := "SELECT COUNT(*) " + baseQuery

	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	dataQuery := `
		SELECT id, name, description, type, price, stock, status, company_id
	` + baseQuery + fmt.Sprintf(`
		ORDER BY id DESC
		LIMIT $%d OFFSET $%d
	`, i, i+1)

	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Type,
			&p.Price,
			&p.Stock,
			&p.Status,
			&p.CompanyID,
		); err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	return products, total, nil
}
