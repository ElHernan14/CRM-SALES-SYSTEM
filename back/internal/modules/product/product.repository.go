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
	GetByID(ctx context.Context, id int) (*models.Product, error)
	Update(ctx context.Context, p *models.Product) error
	SoftDelete(ctx context.Context, id int) error
	GetByIDForUpdate(ctx context.Context, tx *sql.Tx, id int) (*models.Product, error)
	IncrementReservedStock(ctx context.Context, tx *sql.Tx, productID int, quantity int) error
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

func (r *productRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {

	query := `
		SELECT id, name, description, type, price, stock, status, company_id
		FROM product
		WHERE id = $1
		  AND deleted_at IS NULL
		  AND status = 1
	`

	var p models.Product

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Type,
		&p.Price,
		&p.Stock,
		&p.Status,
		&p.CompanyID,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepository) Update(ctx context.Context, p *models.Product) error {

	query := `
		UPDATE product SET
			name = $1,
			description = $2,
			type = $3,
			price = $4,
			stock = $5,
			status = $6
		WHERE id = $7
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		p.Name,
		p.Description,
		p.Type,
		p.Price,
		p.Stock,
		p.Status,
		p.ID,
	)

	return err
}

func (r *productRepository) SoftDelete(ctx context.Context, id int) error {

	query := `
		UPDATE product
		SET 
			deleted_at = NOW(),
			status = 0
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *productRepository) GetByIDForUpdate(
	ctx context.Context,
	tx *sql.Tx,
	id int,
) (*models.Product, error) {

	query := `
		SELECT
			id,
			name,
			price,
			stock,
			reserved_stock,
			company_id
		FROM product
		WHERE id = $1
		FOR UPDATE
	`

	product := &models.Product{}

	err := tx.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.ReservedStock,
		&product.CompanyID,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepository) IncrementReservedStock(
	ctx context.Context,
	tx *sql.Tx,
	productID int,
	quantity int,
) error {

	query := `
		UPDATE product
		SET reserved_stock = reserved_stock + $1
		WHERE id = $2
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		quantity,
		productID,
	)

	return err
}
