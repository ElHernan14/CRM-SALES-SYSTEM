package company

import (
	"context"
	company "crm-system-sales/internal/models/company"
	"database/sql"
	"fmt"
)

type CompanyRepository interface {
	Create(tx *sql.Tx, c *company.Company) (*company.Company, error)
	GetByID(ctx context.Context, id int) (*company.Company, error)
	GetAll(ctx context.Context, search string, limit int, offset int) ([]company.Company, int, error)
	UpdateLogo(ctx context.Context, companyID int, logoPath string) error
}

type companyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) Create(tx *sql.Tx, c *company.Company) (*company.Company, error) {
	query := `
		INSERT INTO company (name, description, status)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := tx.QueryRow(query, c.Name, c.Description, c.Status).Scan(&c.ID, &c.CreatedAt)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *companyRepository) GetByID(ctx context.Context, id int) (*company.Company, error) {
	query := `
		SELECT id, name, description, logo_path, created_at, deleted_at, status
		FROM company
		WHERE id = $1
		  AND deleted_at IS NULL
		  AND status = 1
	`

	var c company.Company
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.Description, &c.LogoPath, &c.CreatedAt, &c.DeletedAt, &c.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *companyRepository) GetAll(ctx context.Context, search string, limit int, offset int) ([]company.Company, int, error) {
	baseQuery := `
		FROM company
		WHERE deleted_at IS NULL
		  AND status = 1
	`

	var args []interface{}
	i := 1

	if search != "" {
		baseQuery += fmt.Sprintf(` AND LOWER(name) LIKE LOWER($%d)`, i)
		args = append(args, "%"+search+"%")
		i++
	}

	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	dataQuery := `SELECT id, name, logo_path ` + baseQuery + fmt.Sprintf(`
		ORDER BY id DESC
		LIMIT $%d OFFSET $%d
	`, i, i+1)

	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var companies []company.Company
	for rows.Next() {
		var c company.Company
		if err := rows.Scan(&c.ID, &c.Name, &c.LogoPath); err != nil {
			return nil, 0, err
		}
		companies = append(companies, c)
	}

	return companies, total, rows.Err()
}

func (r *companyRepository) UpdateLogo(ctx context.Context, companyID int, logoPath string) error {
	query := `
		UPDATE company
		SET logo_path = $1
		WHERE id = $2
		  AND deleted_at IS NULL
		  AND status = 1
	`

	result, err := r.db.ExecContext(ctx, query, logoPath, companyID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
