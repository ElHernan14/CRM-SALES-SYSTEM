package client

import (
	"context"
	"crm-system-sales/internal/core/utils"
	client "crm-system-sales/internal/models/client"
	"fmt"

	"database/sql"
)

type ClientRepository interface {
	Create(tx *sql.Tx, client *client.Client) error
	EmailExists(email string) (bool, error)
	GetClients(ctx context.Context, companyID *int, search string, email string, limit int, offset int) ([]*client.Client, int, error)
}

type clientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) ClientRepository {
	return &clientRepository{db: db}
}

func (r *clientRepository) EmailExists(email string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM client WHERE email = $1)`
	err := r.db.QueryRow(query, email).Scan(&exists)

	return exists, err
}

func (r *clientRepository) Create(tx *sql.Tx, client *client.Client) error {
	query := `
	INSERT INTO client (user_id, company_id, first_name, last_name, email, phone)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id
	`

	return tx.QueryRow(
		query,
		client.UserID,
		client.CompanyID,
		client.FirstName,
		client.LastName,
		client.Email,
		client.Phone,
	).Scan(&client.ID)
}

func (r *clientRepository) GetClients(
	ctx context.Context,
	companyID *int,
	search string,
	email string,
	limit int,
	offset int,
) ([]*client.Client, int, error) {

	var err error
	defer func() {
		utils.Trace(ctx, "REPO GetClients")(err)
	}()

	baseQuery := `
		FROM client
		WHERE 1=1
	`

	args := []interface{}{}
	i := 1

	if companyID != nil {
		baseQuery += fmt.Sprintf(" AND company_id = $%d", i)
		args = append(args, *companyID)
		i++
	}

	if search != "" {
		baseQuery += fmt.Sprintf(`
			AND LOWER(first_name || ' ' || last_name) LIKE LOWER($%d)
		`, i)
		args = append(args, "%"+search+"%")
		i++
	}

	if email != "" {
		baseQuery += fmt.Sprintf(" AND LOWER(email) LIKE LOWER($%d)", i)
		args = append(args, "%"+email+"%")
		i++
	}

	countQuery := "SELECT COUNT(*) " + baseQuery

	var total int
	err = r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	dataQuery := `
		SELECT id, first_name, last_name, email, company_id
	` + baseQuery + `
		ORDER BY last_name DESC
		LIMIT $` + fmt.Sprint(i) + ` OFFSET $` + fmt.Sprint(i+1)

	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var clients []*client.Client

	for rows.Next() {
		var c client.Client

		if err := rows.Scan(
			&c.ID,
			&c.FirstName,
			&c.LastName,
			&c.Email,
			&c.CompanyID,
		); err != nil {
			return nil, 0, err
		}

		clients = append(clients, &c)
	}

	return clients, total, nil
}
