package producttype

import (
	"context"
	"database/sql"
	"fmt"

	model "crm-system-sales/internal/models/product_type"
)

type ProductTypeRepository interface {
	GetByID(ctx context.Context, id int) (*model.ProductType, error)
	GetAll(ctx context.Context, categoryID *int) ([]model.ProductType, error)
}

type productTypeRepository struct {
	db *sql.DB
}

func NewProductTypeRepository(db *sql.DB) ProductTypeRepository {
	return &productTypeRepository{db: db}
}

func (r *productTypeRepository) GetByID(ctx context.Context, id int) (*model.ProductType, error) {
	query := `
		SELECT pt.id, pt.category_id, cp.name, pt.name, COALESCE(pt.description, '')
		FROM product_type pt
		INNER JOIN category_product cp ON cp.id = pt.category_id
		WHERE pt.id = $1
	`

	var t model.ProductType
	err := r.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.CategoryID, &t.Category, &t.Name, &t.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *productTypeRepository) GetAll(ctx context.Context, categoryID *int) ([]model.ProductType, error) {
	query := `
		SELECT pt.id, pt.category_id, cp.name, pt.name, COALESCE(pt.description, '')
		FROM product_type pt
		INNER JOIN category_product cp ON cp.id = pt.category_id
		WHERE true = true
	`

	args := []interface{}{}
	if categoryID != nil {
		query += fmt.Sprintf(" AND pt.category_id = $%d", len(args)+1)
		args = append(args, *categoryID)
	}

	query += " ORDER BY cp.name ASC, pt.name ASC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	types := make([]model.ProductType, 0)
	for rows.Next() {
		var t model.ProductType
		if err := rows.Scan(&t.ID, &t.CategoryID, &t.Category, &t.Name, &t.Description); err != nil {
			return nil, err
		}
		types = append(types, t)
	}

	return types, rows.Err()
}
