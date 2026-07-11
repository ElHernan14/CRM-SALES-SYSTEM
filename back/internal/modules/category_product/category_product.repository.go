package categoryproduct

import (
	"context"
	"database/sql"

	model "crm-system-sales/internal/models/category_product"
)

type CategoryProductRepository interface {
	GetByID(ctx context.Context, id int) (*model.CategoryProduct, error)
	GetAll(ctx context.Context) ([]model.CategoryProduct, error)
}

type categoryProductRepository struct {
	db *sql.DB
}

func NewCategoryProductRepository(db *sql.DB) CategoryProductRepository {
	return &categoryProductRepository{db: db}
}

func (r *categoryProductRepository) GetByID(ctx context.Context, id int) (*model.CategoryProduct, error) {
	query := `
		SELECT id, name, COALESCE(description, '')
		FROM category_product
		WHERE id = $1
	`

	var c model.CategoryProduct
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *categoryProductRepository) GetAll(ctx context.Context) ([]model.CategoryProduct, error) {
	query := `
		SELECT id, name, COALESCE(description, '')
		FROM category_product
		ORDER BY name ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]model.CategoryProduct, 0)
	for rows.Next() {
		var c model.CategoryProduct
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, rows.Err()
}
