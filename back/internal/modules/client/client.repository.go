package client

import (
	"context"
	"crm-system-sales/internal/core/utils"
	client "crm-system-sales/internal/models/client"
	dto "crm-system-sales/internal/modules/client/dto"
	"fmt"
	"strings"

	"database/sql"
)

type ClientRepository interface {
	Create(tx *sql.Tx, client *client.Client) error
	EmailExists(email string) (bool, error)
	GetClients(ctx context.Context, companyID *int, search string, email string, limit int, offset int) ([]*client.Client, int, error)
	GetByID(ctx context.Context, id int) (*client.Client, error)
	Update(ctx context.Context, c *client.Client) error
	SoftDeleteTx(tx *sql.Tx, clientID int) error
	ListCompanyCustomers(
		ctx context.Context,
		companyID int,
		req *dto.GetCompanyCustomersRequest,
	) ([]*client.Client, int, error)
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
		SELECT id, first_name, last_name, email, company_id, phone, status, deleted_at
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
			&c.Phone,
			&c.Status,
			&c.DeletedAt,
		); err != nil {
			return nil, 0, err
		}

		clients = append(clients, &c)
	}

	return clients, total, nil
}

func (r *clientRepository) GetByID(ctx context.Context, id int) (*client.Client, error) {

	var err error
	defer func() {
		utils.Trace(ctx, "REPO GetClientByID")(err)
	}()

	query := `
		SELECT id, first_name, last_name, email, user_id, company_id, phone, status, deleted_at
		FROM client
		WHERE id = $1
	`

	var c client.Client

	err = r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID,
		&c.FirstName,
		&c.LastName,
		&c.Email,
		&c.UserID,
		&c.CompanyID,
		&c.Phone,
		&c.Status,
		&c.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &c, nil
}

func (r *clientRepository) Update(ctx context.Context, c *client.Client) error {

	query := `
		UPDATE client
		SET first_name = $1,
			last_name = $2,
			email = $3,
			phone = $4
		WHERE id = $5
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		c.FirstName,
		c.LastName,
		c.Email,
		c.Phone,
		c.ID,
	)

	return err
}

func (r *clientRepository) SoftDeleteTx(tx *sql.Tx, clientID int) error {

	query := `
		UPDATE client
		SET deleted_at = NOW(),
		    status = 0
		WHERE id = $1
	`

	_, err := tx.Exec(query, clientID)
	return err
}

func (r *clientRepository) ListCompanyCustomers(
	ctx context.Context,
	companyID int,
	req *dto.GetCompanyCustomersRequest,
) ([]*client.Client, int, error) {
	var err error

	allowedSortColumns := map[string]string{
		"name":            "name_complete",
		"total_invoices":  "total_invoices",
		"total_purchased": "total_purchased",
	}

	allowedOrders := map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}

	selectQuery := `
        SELECT
            c.id,
            (c.first_name || ' ' || c.last_name) AS name_complete,
            c.email,
            c.company_id,
            COUNT(i.id) AS total_invoices,
            COALESCE(SUM(i.total_amount), 0) AS total_purchased
    `

	baseQuery := `
        FROM client c
        INNER JOIN invoice i ON i.buyer_client_id = c.id
        WHERE i.seller_company_id = $1
    `

	args := []interface{}{companyID}
	argPos := 2

	// Filtro por nombre (concatenado first_name + last_name)
	if req.Name != "" {
		baseQuery += fmt.Sprintf(
			` AND (
                c.first_name ILIKE $%d
                OR c.last_name ILIKE $%d
            )`,
			argPos, argPos,
		)
		args = append(args, "%"+req.Name+"%")
		argPos++
	}

	// Agrupación
	groupBy := `
        GROUP BY
            c.id,
            c.first_name,
            c.last_name,
            c.email,
            c.company_id
    `

	// Count
	countQuery := "SELECT COUNT(*) FROM (" + selectQuery + baseQuery + groupBy + ") AS sub"
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Ordenamiento
	sortColumn := "name_complete"
	order := "ASC"

	if v, ok := allowedSortColumns[req.SortColumn]; ok {
		sortColumn = v
	}
	if v, ok := allowedOrders[strings.ToLower(req.Order)]; ok {
		order = v
	}

	offset := (req.Page - 1) * req.Limit

	baseQuery += groupBy
	baseQuery += fmt.Sprintf(" ORDER BY %s %s", sortColumn, order)
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)

	args = append(args, req.Limit, offset)

	rows, err := r.db.QueryContext(ctx, selectQuery+baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	customers := make([]*client.Client, 0)
	for rows.Next() {
		var c client.Client
		err := rows.Scan(
			&c.ID,
			&c.NameComplete,
			&c.Email,
			&c.CompanyID,
			&c.TotalInvoices,
			&c.TotalPurchased,
		)
		if err != nil {
			return nil, 0, err
		}
		customers = append(customers, &c)
	}

	return customers, total, nil
}
