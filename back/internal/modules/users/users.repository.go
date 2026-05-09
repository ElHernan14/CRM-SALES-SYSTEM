package users

import (
	auth "crm-system-sales/internal/models/auth"
	users "crm-system-sales/internal/models/users"
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type UserRepository interface {
	GetUserLogin(email string) (*auth.UserLogin, error)
	EmailExists(email string) (bool, error)
	GetByID(id int) (*users.User, error)
	GetAll() ([]users.User, error)
	Create(tx *sql.Tx, user *users.User) (*users.User, error)
	SoftDeleteByIDTx(tx *sql.Tx, userID int) error
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetUserLogin(email string) (*auth.UserLogin, error) {
	query := `
		SELECT 
			u.id,
			u.email,
			u.password_hash,
			c.company_id,
			c.client_id,
			COALESCE(ARRAY_AGG(DISTINCT r.name) FILTER (WHERE r.name IS NOT NULL), '{}') AS roles,
			COALESCE(ARRAY_AGG(DISTINCT p.name) FILTER (WHERE p.name IS NOT NULL), '{}') AS permissions
		FROM users u
		LEFT JOIN client c ON c.user_id = u.id
		LEFT JOIN user_rol ur ON ur.user_id = u.id
		LEFT JOIN rol r ON r.id = ur.role_id
		LEFT JOIN role_permission rp ON rp.role_id = r.id
		LEFT JOIN permission p ON p.id = rp.permission_id
		WHERE u.email = $1 AND u.status = 1
		GROUP BY u.id, u.email, u.password_hash, c.company_id, c.client_id;
	`

	row := r.DB.QueryRow(query, email)

	var user auth.UserLogin
	var companyID sql.NullInt64
	var clientID sql.NullInt64

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&companyID,
		&clientID,
		pq.Array(&user.Roles),
		pq.Array(&user.Permissions),
	)

	if err != nil {
		log.Println("Error scanning user data: ", err)
		return nil, err
	}

	if companyID.Valid {
		id := int(companyID.Int64)
		user.CompanyID = &id
	}

	if clientID.Valid {
		id := int(clientID.Int64)
		user.ClientID = &id
	}

	return &user, nil
}

func (r *userRepository) Create(tx *sql.Tx, user *users.User) (*users.User, error) {
	err := tx.QueryRow(`
        INSERT INTO users (email, password_hash)
        VALUES ($1, $2)
        RETURNING id
    `, user.Email, user.PasswordHash).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) EmailExists(email string) (bool, error) {
	var exists bool
	err := r.DB.QueryRow(`
        SELECT EXISTS (
            SELECT 1 FROM users WHERE email = $1
        )
    `, email).Scan(&exists)
	return exists, err
}

func (r *userRepository) GetByID(id int) (*users.User, error) {
	return nil, nil
}

func (r *userRepository) GetAll() ([]users.User, error) {
	return nil, nil
}

func (r *userRepository) SoftDeleteByIDTx(tx *sql.Tx, userID int) error {

	query := `
		UPDATE users
		SET status = 0,
			deleted_at = NOW()
		WHERE id = $1
	`

	_, err := tx.Exec(query, userID)
	return err
}
