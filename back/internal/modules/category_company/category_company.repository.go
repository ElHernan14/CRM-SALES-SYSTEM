package categorycompany

import (
	"context"
	"database/sql"

	model "crm-system-sales/internal/models/category_company"
)

type CategoryCompanyRepository interface {
	GetByID(ctx context.Context, id int) (*model.CategoryCompany, error)
	GetAll(ctx context.Context) ([]model.CategoryCompany, error)
}

type categoryCompanyRepository struct {
	db *sql.DB
}

func NewCategoryCompanyRepository(db *sql.DB) CategoryCompanyRepository {
	return &categoryCompanyRepository{db: db}
}

func (r *categoryCompanyRepository) GetByID(ctx context.Context, id int) (*model.CategoryCompany, error) {
	query := `
		SELECT id, name, description
		FROM category_company
		WHERE id = $1
	`

	var c model.CategoryCompany
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *categoryCompanyRepository) GetAll(ctx context.Context) ([]model.CategoryCompany, error) {
	query := `
		SELECT id, name, description
		FROM category_company
		ORDER BY name ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]model.CategoryCompany, 0)
	for rows.Next() {
		var c model.CategoryCompany
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, rows.Err()
}
