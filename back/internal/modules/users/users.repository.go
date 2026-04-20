package users

import (
	models "crm-system-sales/internal/models/auth"
	"database/sql"

	"github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserLogin(email string) (*models.UserLogin, error) {
	query := `
		SELECT 
			u.id,
			u.email,
			u.password_hash,
			c.company_id,
			COALESCE(ARRAY_AGG(DISTINCT p.name) FILTER (WHERE p.name IS NOT NULL), '{}') AS permissions
		FROM users u
		LEFT JOIN clients c ON c.user_id = u.id
		LEFT JOIN user_rol ur ON ur.user_id = u.id
		LEFT JOIN rol r ON r.id = ur.role_id
		LEFT JOIN role_permission rp ON rp.role_id = r.id
		LEFT JOIN permission p ON p.id = rp.permission_id
		WHERE u.email = $1
		GROUP BY u.id, u.email, u.password_hash, c.company_id;
	`

	row := r.DB.QueryRow(query, email)

	var user models.UserLogin
	var companyID sql.NullInt64

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&companyID,
		pq.Array(&user.Permissions),
	)

	if err != nil {
		return nil, err
	}

	if companyID.Valid {
		id := int(companyID.Int64)
		user.CompanyID = &id
	}

	return &user, nil
}
