package invoiceRepository

import (
	"context"
	tenantHelper "crm-system-sales/internal/core/tenant"
	invoicedto "crm-system-sales/internal/modules/invoice/dto"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"
	storedto "crm-system-sales/internal/modules/store/dto"
	"database/sql"
	"fmt"
	"strings"
)

type InvoiceRepository interface {
	Create(ctx context.Context, tx *sql.Tx, invoice *invoiceModel.Invoice) error
	GetByID(ctx context.Context, id int) (*invoiceModel.Invoice, error)
	RecalculateInvoiceTotals(ctx context.Context, tx *sql.Tx, invoiceID int) error
	GetActiveDraft(ctx context.Context, buyerClientID int, sellerCompanyID int, source string) (*invoiceModel.Invoice, error)
	UpdateStatus(ctx context.Context, tx *sql.Tx, invoiceID int, status string) error
	SetPaidAmount(ctx context.Context, tx *sql.Tx, invoiceID int, paidAmount float64) error
	ListBySellerCompanyID(ctx context.Context, companyID int, req *invoicedto.GetCompanyInvoicesRequest) ([]*invoiceModel.Invoice, int, error)
	ListByBuyerCompanyID(ctx context.Context, companyID int, req *invoicedto.GetCompanyInvoicesRequest) ([]*invoiceModel.Invoice, int, error)
	ListActiveDraftsByBuyer(ctx context.Context, buyerClientID int, invoiceIDs []int) ([]*invoiceModel.Invoice, error)
	ListStorePurchasesByBuyer(ctx context.Context, buyerClientID int, req *storedto.GetStorePurchasesRequest) ([]*invoiceModel.Invoice, int, error)
	GetActiveDraftByTenant(
		ctx context.Context,
		tenant *tenantHelper.TenantContext,
	) (*invoiceModel.Invoice, error)
}

type invoiceRepository struct {
	DB *sql.DB
}

func NewInvoiceRepository(db *sql.DB) InvoiceRepository {
	return &invoiceRepository{DB: db}
}

const DefaultTaxRate = 0.21 // 21%

func (r *invoiceRepository) Create(
	ctx context.Context,
	tx *sql.Tx,
	invoice *invoiceModel.Invoice,
) error {

	query := `
		INSERT INTO invoice (
			buyer_client_id,
			seller_company_id,
			created_by_user_id,
			total_amount,
			status_invoice,
			source
		)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id, created_at
	`
	if invoice.Source == "" {
		invoice.Source = "erp"
	}
	if tx != nil {
		return tx.QueryRowContext(
			ctx,
			query,
			invoice.BuyerClientID,
			invoice.SellerCompanyID,
			invoice.CreatedByUserID,
			invoice.TotalAmount,
			invoice.StatusInvoice,
			invoice.Source,
		).Scan(
			&invoice.ID,
			&invoice.CreatedAt,
		)
	} else {
		return r.DB.QueryRowContext(
			ctx,
			query,
			invoice.BuyerClientID,
			invoice.SellerCompanyID,
			invoice.CreatedByUserID,
			invoice.TotalAmount,
			invoice.StatusInvoice,
			invoice.Source,
		).Scan(
			&invoice.ID,
			&invoice.CreatedAt,
		)
	}
}

func (r *invoiceRepository) GetByID(ctx context.Context, id int) (*invoiceModel.Invoice, error) {
	var err error

	query := `
        SELECT id, buyer_client_id, seller_company_id, created_by_user_id,
               total_amount, subtotal, paid_amount, status_invoice, COALESCE(source, 'erp') AS source, status, taxes, created_at, updated_at, deleted_at
        FROM invoice
        WHERE id = $1
    `

	var inv invoiceModel.Invoice

	err = r.DB.QueryRowContext(ctx, query, id).Scan(
		&inv.ID,
		&inv.BuyerClientID,
		&inv.SellerCompanyID,
		&inv.CreatedByUserID,
		&inv.TotalAmount,
		&inv.Subtotal,
		&inv.PaidAmount,
		&inv.StatusInvoice,
		&inv.Source,
		&inv.Status,
		&inv.Taxes,
		&inv.CreatedAt,
		&inv.UpdatedAt,
		&inv.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &inv, nil
}

func (r *invoiceRepository) RecalculateInvoiceTotals(
	ctx context.Context,
	tx *sql.Tx,
	invoiceID int,
) error {

	//  subtotal items
	var subtotal float64

	querySubtotal := `
		SELECT COALESCE(SUM(subtotal), 0)
		FROM invoice_item
		WHERE invoice_id = $1
		  AND status = 1
		  AND deleted_at IS NULL
	`

	err := tx.QueryRowContext(
		ctx,
		querySubtotal,
		invoiceID,
	).Scan(&subtotal)

	if err != nil {
		return err
	}

	//  taxes
	taxes := subtotal * DefaultTaxRate

	//  total
	total := subtotal + taxes

	//  update invoice
	queryUpdate := `
		UPDATE invoice
		SET
			subtotal = $1,
			taxes = $2,
			total_amount = $3,
			updated_at = NOW()
		WHERE id = $4
	`

	_, err = tx.ExecContext(
		ctx,
		queryUpdate,
		subtotal,
		taxes,
		total,
		invoiceID,
	)

	return err
}

func (r *invoiceRepository) GetActiveDraft(
	ctx context.Context,
	buyerClientID int,
	sellerCompanyID int,
	source string,
) (*invoiceModel.Invoice, error) {

	query := `
		SELECT
			id,
			buyer_client_id,
			seller_company_id,
			status_invoice,
			COALESCE(source, 'erp') AS source,
			created_at,
			created_by_user_id,
			total_amount,
			subtotal,
			taxes
		FROM invoice
		WHERE
			buyer_client_id = $1
			AND seller_company_id = $2
			AND COALESCE(source, 'erp') = $3
			AND status_invoice = 'draft'
			AND status = 1
		ORDER BY created_at DESC
		LIMIT 1
	`

	invoice := &invoiceModel.Invoice{}

	err := r.DB.QueryRowContext(
		ctx,
		query,
		buyerClientID,
		sellerCompanyID,
		source,
	).Scan(
		&invoice.ID,
		&invoice.BuyerClientID,
		&invoice.SellerCompanyID,
		&invoice.StatusInvoice,
		&invoice.Source,
		&invoice.CreatedAt,
		&invoice.CreatedByUserID,
		&invoice.TotalAmount,
		&invoice.Subtotal,
		&invoice.Taxes,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (r *invoiceRepository) UpdateStatus(
	ctx context.Context,
	tx *sql.Tx,
	invoiceID int,
	status string,
) error {

	query := `
		UPDATE invoice
		SET
			status_invoice = $1,
			updated_at = NOW()
		WHERE id = $2
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		status,
		invoiceID,
	)

	return err
}

func (r *invoiceRepository) SetPaidAmount(
	ctx context.Context,
	tx *sql.Tx,
	invoiceID int,
	paidAmount float64,
) error {

	query := `
		UPDATE invoice
		SET
			paid_amount = $1,
			updated_at = NOW()
		WHERE id = $2
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		paidAmount,
		invoiceID,
	)

	return err
}

func (r *invoiceRepository) ListBySellerCompanyID(
	ctx context.Context,
	companyID int,
	req *invoicedto.GetCompanyInvoicesRequest,
) ([]*invoiceModel.Invoice, int, error) {
	var err error

	allowedSortColumns := map[string]string{
		"created_at":     "created_at",
		"total_amount":   "total_amount",
		"paid_amount":    "paid_amount",
		"status_invoice": "status_invoice",
	}

	allowedOrders := map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}

	selectQuery := `
        SELECT
            i.id,
			i.buyer_client_id,
			i.seller_company_id,
			i.status_invoice,
			COALESCE(i.source, 'erp') AS source,
			i.total_amount,
			i.paid_amount,
			i.created_at,
			cb.first_name AS buyer_first_name,
			cb.last_name AS buyer_last_name,
			bco.name AS buyer_company_name,
			co.name AS seller_company_name
    `

	baseQuery := `
        FROM invoice i
		LEFT JOIN client cb ON cb.id = i.buyer_client_id
		LEFT JOIN company bco ON bco.id = cb.company_id
		LEFT JOIN company co ON co.id = i.seller_company_id
		WHERE i.seller_company_id = $1
    `

	args := []interface{}{companyID}
	argPos := 2

	// Filtro por status_invoice
	if req.StatusInvoice != "" {
		baseQuery += fmt.Sprintf(" AND status_invoice = $%d", argPos)
		args = append(args, req.StatusInvoice)
		argPos++
	}

	if req.BuyerName != "" {
		baseQuery += fmt.Sprintf(" AND (cb.first_name ILIKE $%d OR cb.last_name ILIKE $%d OR CONCAT(cb.first_name, ' ', cb.last_name) ILIKE $%d OR bco.name ILIKE $%d)", argPos, argPos, argPos, argPos)
		args = append(args, "%"+req.BuyerName+"%")
		argPos++
	}

	if req.SellerCompany != "" {
		baseQuery += fmt.Sprintf(" AND co.name ILIKE $%d", argPos)
		args = append(args, "%"+req.SellerCompany+"%")
		argPos++
	}

	// Filtro por status
	if req.Status != 0 {
		baseQuery += fmt.Sprintf(" AND i.status = $%d", argPos)
		args = append(args, req.Status)
		argPos++
	}

	// Count
	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	if err := r.DB.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Ordenamiento
	sortColumn := "created_at"
	order := "DESC"

	if v, ok := allowedSortColumns[req.SortColumn]; ok {
		sortColumn = v
	}
	if v, ok := allowedOrders[strings.ToLower(req.Order)]; ok {
		order = v
	}

	offset := (req.Page - 1) * req.Limit

	baseQuery += fmt.Sprintf(" ORDER BY %s %s", sortColumn, order)
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)

	args = append(args, req.Limit, offset)

	rows, err := r.DB.QueryContext(ctx, selectQuery+baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	invoices := make([]*invoiceModel.Invoice, 0)
	for rows.Next() {
		var inv invoiceModel.Invoice
		var buyerFirstName, buyerLastName, buyerCompanyName sql.NullString
		var sellerCompanyName sql.NullString

		err := rows.Scan(
			&inv.ID,
			&inv.BuyerClientID,
			&inv.SellerCompanyID,
			&inv.StatusInvoice,
			&inv.Source,
			&inv.TotalAmount,
			&inv.PaidAmount,
			&inv.CreatedAt,
			&buyerFirstName,
			&buyerLastName,
			&buyerCompanyName,
			&sellerCompanyName,
		)
		if err != nil {
			return nil, 0, err
		}

		inv.BuyerName = strings.TrimSpace(buyerFirstName.String + " " + buyerLastName.String)
		if buyerCompanyName.Valid && strings.TrimSpace(buyerCompanyName.String) != "" {
			inv.BuyerName = buyerCompanyName.String
		}
		inv.SellerName = sellerCompanyName.String

		invoices = append(invoices, &inv)
	}

	return invoices, total, nil
}

func (r *invoiceRepository) GetActiveDraftByTenant(
	ctx context.Context,
	tenant *tenantHelper.TenantContext,
) (*invoiceModel.Invoice, error) {

	query := `
        SELECT
            id,
            buyer_client_id,
            seller_company_id,
            status_invoice,
			COALESCE(source, 'erp') AS source,
            created_at,
            created_by_user_id,
            total_amount
        FROM invoice
        WHERE
            buyer_client_id = $1
			AND source = 'store'
            AND status_invoice = 'draft'
            AND status = 1
        ORDER BY created_at DESC
        LIMIT 1
    `

	invoice := &invoiceModel.Invoice{}

	err := r.DB.QueryRowContext(
		ctx,
		query,
		*tenant.ClientID,
	).Scan(
		&invoice.ID,
		&invoice.BuyerClientID,
		&invoice.SellerCompanyID,
		&invoice.StatusInvoice,
		&invoice.Source,
		&invoice.CreatedAt,
		&invoice.CreatedByUserID,
		&invoice.TotalAmount,
	)

	if err != nil {
		return nil, err
	}

	return invoice, nil
}
func (r *invoiceRepository) ListByBuyerCompanyID(
	ctx context.Context,
	companyID int,
	req *invoicedto.GetCompanyInvoicesRequest,
) ([]*invoiceModel.Invoice, int, error) {
	allowedSortColumns := map[string]string{
		"created_at":     "i.created_at",
		"total_amount":   "i.total_amount",
		"paid_amount":    "i.paid_amount",
		"status_invoice": "i.status_invoice",
	}

	allowedOrders := map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}

	selectQuery := `
		SELECT
			i.id,
			i.buyer_client_id,
			i.seller_company_id,
			i.status_invoice,
			COALESCE(i.source, 'erp') AS source,
			i.total_amount,
			i.paid_amount,
			i.created_at,
			COALESCE(NULLIF(bco.name, ''), TRIM(cb.first_name || ' ' || cb.last_name)) AS buyer_name,
			co.name AS seller_company_name
	`

	baseQuery := `
		FROM invoice i
		LEFT JOIN client cb ON cb.id = i.buyer_client_id
		LEFT JOIN company bco ON bco.id = cb.company_id
		LEFT JOIN company co ON co.id = i.seller_company_id
		WHERE cb.company_id = $1
		  AND COALESCE(i.source, 'erp') = 'erp'
	`

	args := []interface{}{companyID}
	argPos := 2

	if req.StatusInvoice != "" {
		baseQuery += fmt.Sprintf(" AND i.status_invoice = $%d", argPos)
		args = append(args, req.StatusInvoice)
		argPos++
	}

	if req.BuyerName != "" {
		baseQuery += fmt.Sprintf(" AND (cb.first_name ILIKE $%d OR cb.last_name ILIKE $%d OR CONCAT(cb.first_name, ' ', cb.last_name) ILIKE $%d OR bco.name ILIKE $%d)", argPos, argPos, argPos, argPos)
		args = append(args, "%"+req.BuyerName+"%")
		argPos++
	}

	if req.SellerCompany != "" {
		baseQuery += fmt.Sprintf(" AND co.name ILIKE $%d", argPos)
		args = append(args, "%"+req.SellerCompany+"%")
		argPos++
	}

	if req.Status != 0 {
		baseQuery += fmt.Sprintf(" AND i.status = $%d", argPos)
		args = append(args, req.Status)
		argPos++
	}

	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	if err := r.DB.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	sortColumn := "i.created_at"
	order := "DESC"
	if v, ok := allowedSortColumns[req.SortColumn]; ok {
		sortColumn = v
	}
	if v, ok := allowedOrders[strings.ToLower(req.Order)]; ok {
		order = v
	}

	offset := (req.Page - 1) * req.Limit
	baseQuery += fmt.Sprintf(" ORDER BY %s %s", sortColumn, order)
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, req.Limit, offset)

	rows, err := r.DB.QueryContext(ctx, selectQuery+baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	invoices := make([]*invoiceModel.Invoice, 0)
	for rows.Next() {
		var inv invoiceModel.Invoice
		var buyerName, sellerCompanyName sql.NullString

		if err := rows.Scan(
			&inv.ID,
			&inv.BuyerClientID,
			&inv.SellerCompanyID,
			&inv.StatusInvoice,
			&inv.Source,
			&inv.TotalAmount,
			&inv.PaidAmount,
			&inv.CreatedAt,
			&buyerName,
			&sellerCompanyName,
		); err != nil {
			return nil, 0, err
		}

		inv.BuyerName = buyerName.String
		inv.SellerName = sellerCompanyName.String
		invoices = append(invoices, &inv)
	}

	return invoices, total, rows.Err()
}

func (r *invoiceRepository) ListActiveDraftsByBuyer(ctx context.Context, buyerClientID int, invoiceIDs []int) ([]*invoiceModel.Invoice, error) {
	query := `
		SELECT
			i.id,
			i.buyer_client_id,
			i.seller_company_id,
			i.status_invoice,
			i.source,
			i.subtotal,
			i.taxes,
			i.total_amount,
			co.name AS seller_company_name
		FROM invoice i
		INNER JOIN company co ON co.id = i.seller_company_id
		WHERE i.buyer_client_id = $1
		  AND i.status_invoice = 'draft'
		  AND i.source = 'store'
		  AND i.status = 1
	`

	args := []interface{}{buyerClientID}
	argPos := 2
	if len(invoiceIDs) > 0 {
		placeholders := make([]string, 0, len(invoiceIDs))
		for _, id := range invoiceIDs {
			placeholders = append(placeholders, fmt.Sprintf("$%d", argPos))
			args = append(args, id)
			argPos++
		}
		query += " AND i.id IN (" + strings.Join(placeholders, ",") + ")"
	}
	query += " ORDER BY co.name ASC, i.id ASC"

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoices := make([]*invoiceModel.Invoice, 0)
	for rows.Next() {
		var inv invoiceModel.Invoice
		var sellerCompanyName sql.NullString
		if err := rows.Scan(
			&inv.ID,
			&inv.BuyerClientID,
			&inv.SellerCompanyID,
			&inv.StatusInvoice,
			&inv.Source,
			&inv.Subtotal,
			&inv.Taxes,
			&inv.TotalAmount,
			&sellerCompanyName,
		); err != nil {
			return nil, err
		}
		inv.SellerName = sellerCompanyName.String
		invoices = append(invoices, &inv)
	}

	return invoices, rows.Err()
}

func (r *invoiceRepository) ListStorePurchasesByBuyer(ctx context.Context, buyerClientID int, req *storedto.GetStorePurchasesRequest) ([]*invoiceModel.Invoice, int, error) {
	allowedSortColumns := map[string]string{
		"created_at":     "i.created_at",
		"total_amount":   "i.total_amount",
		"paid_amount":    "i.paid_amount",
		"status_invoice": "i.status_invoice",
	}
	allowedOrders := map[string]string{"asc": "ASC", "desc": "DESC"}

	selectQuery := `
		SELECT
			i.id,
			i.seller_company_id,
			co.name AS seller_company_name,
			i.status_invoice,
			i.source,
			i.subtotal,
			i.taxes,
			i.total_amount,
			i.paid_amount,
			COUNT(ii.id) AS item_count,
			i.created_at
	`
	baseQuery := `
		FROM invoice i
		INNER JOIN company co ON co.id = i.seller_company_id
		LEFT JOIN invoice_item ii ON ii.invoice_id = i.id AND ii.status = 1
		WHERE i.buyer_client_id = $1
		  AND i.status = 1
		  AND i.source = 'store'
		  AND i.status_invoice <> 'draft'
	`
	args := []interface{}{buyerClientID}
	argPos := 2

	if req.StatusInvoice != "" {
		baseQuery += fmt.Sprintf(" AND i.status_invoice = $%d", argPos)
		args = append(args, req.StatusInvoice)
		argPos++
	}
	if req.SellerCompanyID != nil {
		baseQuery += fmt.Sprintf(" AND i.seller_company_id = $%d", argPos)
		args = append(args, *req.SellerCompanyID)
		argPos++
	}
	if req.Search != "" {
		baseQuery += fmt.Sprintf(` AND (
			co.name ILIKE $%d
			OR EXISTS (
				SELECT 1
				FROM invoice_item search_item
				WHERE search_item.invoice_id = i.id
				  AND search_item.status = 1
				  AND search_item.product_name ILIKE $%d
			)
		)`, argPos, argPos)
		args = append(args, "%"+req.Search+"%")
		argPos++
	}

	groupBy := `
		GROUP BY
			i.id,
			i.seller_company_id,
			co.name,
			i.status_invoice,
			i.source,
			i.subtotal,
			i.taxes,
			i.total_amount,
			i.paid_amount,
			i.created_at
	`

	countQuery := "SELECT COUNT(*) FROM (SELECT i.id " + baseQuery + groupBy + ") AS purchases_count"
	var total int
	if err := r.DB.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	sortColumn := "i.created_at"
	order := "DESC"
	if v, ok := allowedSortColumns[req.SortColumn]; ok {
		sortColumn = v
	}
	if v, ok := allowedOrders[strings.ToLower(req.Order)]; ok {
		order = v
	}

	offset := (req.Page - 1) * req.Limit
	dataQuery := selectQuery + baseQuery + groupBy + fmt.Sprintf(" ORDER BY %s %s LIMIT $%d OFFSET $%d", sortColumn, order, argPos, argPos+1)
	args = append(args, req.Limit, offset)

	rows, err := r.DB.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	invoices := make([]*invoiceModel.Invoice, 0)
	for rows.Next() {
		var inv invoiceModel.Invoice
		if err := rows.Scan(
			&inv.ID,
			&inv.SellerCompanyID,
			&inv.SellerName,
			&inv.StatusInvoice,
			&inv.Source,
			&inv.Subtotal,
			&inv.Taxes,
			&inv.TotalAmount,
			&inv.PaidAmount,
			&inv.ItemCount,
			&inv.CreatedAt,
		); err != nil {
			return nil, 0, err
		}
		invoices = append(invoices, &inv)
	}

	return invoices, total, rows.Err()
}
