package auth

import "database/sql"

type AuthRepository interface {
	AssignRole(tx *sql.Tx, userID int, roleName string) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) AssignRole(tx *sql.Tx, userID int, roleName string) error {
	_, err := tx.Exec(`
        INSERT INTO user_rol (user_id, role_id)
        VALUES ($1, (SELECT id FROM rol WHERE name = $2))
    `, userID, roleName)

	return err
}
