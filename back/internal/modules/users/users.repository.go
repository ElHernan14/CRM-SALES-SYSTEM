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

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT 
			u.id,
			u.email,
			u.password_hash,
			ARRAY_AGG(DISTINCT r.name) AS roles,
			ARRAY_AGG(DISTINCT p.name) AS permissions
		FROM users u
		LEFT JOIN user_rol ur ON ur.user_id = u.id
		LEFT JOIN rol r ON r.id = ur.role_id
		LEFT JOIN role_permission rp ON rp.role_id = r.id
		LEFT JOIN permission p ON p.id = rp.permission_id
		WHERE u.email = $1
		GROUP BY u.id, u.email, u.password_hash;

	`

	row := r.DB.QueryRow(query, email)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		pq.Array(&user.Permissions),
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
