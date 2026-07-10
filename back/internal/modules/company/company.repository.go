package company

import (
	"context"
	"database/sql"
	"fmt"

	company "crm-system-sales/internal/models/company"
)

type CompanyRepository interface {
	Create(tx *sql.Tx, c *company.Company) (*company.Company, error)
	GetByID(ctx context.Context, id int) (*company.Company, error)
	GetAll(ctx context.Context, search string, categoryID *int, category string, limit int, offset int) ([]company.Company, int, error)
	UpdateLogo(ctx context.Context, companyID int, logoPath string) error
	UpdateCoverImage(ctx context.Context, companyID int, coverImagePath string) error
}

type companyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) Create(tx *sql.Tx, c *company.Company) (*company.Company, error) {
	query := `
		INSERT INTO company (name, category_id, description, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	err := tx.QueryRow(query, c.Name, c.CategoryID, c.Description, c.Status).Scan(&c.ID, &c.CreatedAt)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *companyRepository) GetByID(ctx context.Context, id int) (*company.Company, error) {
	query := `
		SELECT c.id, c.name, c.category_id, cc.name, c.description, c.logo_path, c.cover_image_path, c.created_at, c.deleted_at, c.status
		FROM company c
		INNER JOIN category_company cc ON cc.id = c.category_id
		WHERE c.id = $1
		  AND c.deleted_at IS NULL
		  AND c.status = 1
	`

	var c company.Company
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.CategoryID, &c.Category, &c.Description, &c.LogoPath, &c.CoverImagePath, &c.CreatedAt, &c.DeletedAt, &c.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *companyRepository) GetAll(ctx context.Context, search string, categoryID *int, category string, limit int, offset int) ([]company.Company, int, error) {
	baseQuery := `
		FROM company c
		INNER JOIN category_company cc ON cc.id = c.category_id
		WHERE c.deleted_at IS NULL
		  AND c.status = 1
	`

	var args []interface{}
	i := 1

	if search != "" {
		baseQuery += fmt.Sprintf(` AND LOWER(c.name) LIKE LOWER($%d)`, i)
		args = append(args, "%"+search+"%")
		i++
	}

	if categoryID != nil {
		baseQuery += fmt.Sprintf(` AND c.category_id = $%d`, i)
		args = append(args, *categoryID)
		i++
	} else if category != "" {
		baseQuery += fmt.Sprintf(` AND LOWER(cc.name) = LOWER($%d)`, i)
		args = append(args, category)
		i++
	}

	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	dataQuery := `SELECT c.id, c.name, c.category_id, cc.name, c.logo_path, c.cover_image_path ` + baseQuery + fmt.Sprintf(`
		ORDER BY c.id DESC
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
		if err := rows.Scan(&c.ID, &c.Name, &c.CategoryID, &c.Category, &c.LogoPath, &c.CoverImagePath); err != nil {
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

	return r.updateImageField(ctx, query, logoPath, companyID)
}

func (r *companyRepository) UpdateCoverImage(ctx context.Context, companyID int, coverImagePath string) error {
	query := `
		UPDATE company
		SET cover_image_path = $1
		WHERE id = $2
		  AND deleted_at IS NULL
		  AND status = 1
	`

	return r.updateImageField(ctx, query, coverImagePath, companyID)
}

func (r *companyRepository) updateImageField(ctx context.Context, query string, path string, companyID int) error {
	result, err := r.db.ExecContext(ctx, query, path, companyID)
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
