package product

import (
	"context"
	models "crm-system-sales/internal/models/product"
	productdto "crm-system-sales/internal/modules/product/dto"
	storedto "crm-system-sales/internal/modules/store/dto"
	"database/sql"
	"fmt"
	"strings"
)

type ProductRepository interface {
	Create(ctx context.Context, p *models.Product) error
	GetAll(ctx context.Context, search string, productType string, minPrice float64, maxPrice float64, companyID *int, limit int, offset int) ([]models.Product, int, error)
	GetByID(ctx context.Context, id int) (*models.Product, error)
	Update(ctx context.Context, p *models.Product) error
	UpdateImage(ctx context.Context, productID int, imagePath string) error
	SoftDelete(ctx context.Context, id int) error
	GetByIDForUpdate(ctx context.Context, tx *sql.Tx, id int) (*models.Product, error)
	ListByCompanyID(ctx context.Context, companyID int, req *productdto.GetCompanyProductsRequest) ([]*models.Product, int, error)
	ListAvailableProducts(ctx context.Context, req *storedto.GetStoreProductsRequest) ([]*models.Product, int, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, p *models.Product) error {
	query := `
		INSERT INTO product (company_id, name, description, type, price, stock, image_path, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, 1)
		RETURNING id
	`

	return r.db.QueryRowContext(ctx, query, p.CompanyID, p.Name, p.Description, p.Type, p.Price, p.Stock, p.ImagePath).Scan(&p.ID)
}

func (r *productRepository) GetAll(ctx context.Context, search string, productType string, minPrice float64, maxPrice float64, companyID *int, limit int, offset int) ([]models.Product, int, error) {
	baseQuery := `
		FROM product
		WHERE true = true
	`

	var args []interface{}
	i := 1

	if companyID != nil {
		baseQuery += fmt.Sprintf(" AND company_id = $%d", i)
		args = append(args, *companyID)
		i++
	}

	if search != "" {
		baseQuery += fmt.Sprintf(` AND LOWER(name) LIKE LOWER($%d)`, i)
		args = append(args, "%"+search+"%")
		i++
	}

	if productType != "" {
		baseQuery += fmt.Sprintf(" AND type = $%d", i)
		args = append(args, productType)
		i++
	}

	if minPrice > 0 {
		baseQuery += fmt.Sprintf(" AND price >= $%d", i)
		args = append(args, minPrice)
		i++
	}

	if maxPrice > 0 {
		baseQuery += fmt.Sprintf(" AND price <= $%d", i)
		args = append(args, maxPrice)
		i++
	}

	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	dataQuery := `SELECT id, name, description, type, price, stock, status, company_id, reserved_stock, image_path ` + baseQuery + fmt.Sprintf(`
		ORDER BY id DESC
		LIMIT $%d OFFSET $%d
	`, i, i+1)

	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Type, &p.Price, &p.Stock, &p.Status, &p.CompanyID, &p.ReservedStock, &p.ImagePath); err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	return products, total, rows.Err()
}

func (r *productRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	query := `
		SELECT id, name, description, type, price, stock, status, company_id, image_path
		FROM product
		WHERE id = $1
	`

	var p models.Product
	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Type, &p.Price, &p.Stock, &p.Status, &p.CompanyID, &p.ImagePath)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepository) Update(ctx context.Context, p *models.Product) error {
	query := `
		UPDATE product SET
			name = $1,
			description = $2,
			type = $3,
			price = $4,
			stock = $5,
			status = $6
		WHERE id = $7
	`

	_, err := r.db.ExecContext(ctx, query, p.Name, p.Description, p.Type, p.Price, p.Stock, p.Status, p.ID)
	return err
}

func (r *productRepository) UpdateImage(ctx context.Context, productID int, imagePath string) error {
	query := `
		UPDATE product
		SET image_path = $1
		WHERE id = $2
		  AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, imagePath, productID)
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

func (r *productRepository) SoftDelete(ctx context.Context, id int) error {
	query := `
		UPDATE product
		SET deleted_at = NOW(), status = 0
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *productRepository) GetByIDForUpdate(ctx context.Context, tx *sql.Tx, id int) (*models.Product, error) {
	query := `
		SELECT id, name, price, stock, reserved_stock, company_id
		FROM product
		WHERE id = $1
		FOR UPDATE
	`

	product := &models.Product{}
	err := tx.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.ReservedStock, &product.CompanyID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepository) ListByCompanyID(ctx context.Context, companyID int, req *productdto.GetCompanyProductsRequest) ([]*models.Product, int, error) {
	allowedSortColumns := map[string]string{"name": "name", "description": "description", "type": "type", "price": "price", "stock": "stock"}
	allowedOrders := map[string]string{"asc": "ASC", "desc": "DESC"}

	selectQuery := `SELECT id, name, description, type, price, stock, reserved_stock, status, image_path `
	baseQuery := `FROM product WHERE company_id = $1`
	args := []interface{}{companyID}
	argPos := 2

	if req.Name != "" {
		baseQuery += fmt.Sprintf(" AND name ILIKE $%d", argPos)
		args = append(args, "%"+req.Name+"%")
		argPos++
	}
	if req.Type != "" {
		baseQuery += fmt.Sprintf(" AND type = $%d", argPos)
		args = append(args, req.Type)
		argPos++
	}
	if req.Status != 0 {
		baseQuery += fmt.Sprintf(" AND status = $%d", argPos)
		args = append(args, req.Status)
		argPos++
	}

	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

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

	rows, err := r.db.QueryContext(ctx, selectQuery+baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	products := make([]*models.Product, 0)
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Type, &p.Price, &p.Stock, &p.ReservedStock, &p.Status, &p.ImagePath); err != nil {
			return nil, 0, err
		}
		products = append(products, &p)
	}

	return products, total, rows.Err()
}

func (r *productRepository) ListAvailableProducts(ctx context.Context, req *storedto.GetStoreProductsRequest) ([]*models.Product, int, error) {
	allowedSortColumns := map[string]string{"name": "p.name", "created_at": "p.created_at", "type": "p.type", "price": "p.price", "stock": "p.stock"}
	allowedOrders := map[string]string{"asc": "ASC", "desc": "DESC"}

	selectQuery := `
		SELECT
			p.id,
			p.company_id,
			c.name AS company_name,
			p.name,
			p.description,
			p.type,
			p.price,
			p.stock,
			p.reserved_stock,
			p.image_path
	`

	baseQuery := `
		FROM product p
		INNER JOIN company c ON c.id = p.company_id
		WHERE p.status = 1
		  AND p.deleted_at IS NULL
		  AND c.status = 1
		  AND c.deleted_at IS NULL
		  AND p.stock > p.reserved_stock
	`

	args := []interface{}{}
	argPos := 1

	if req.CompanyID != nil {
		baseQuery += fmt.Sprintf(" AND p.company_id = $%d", argPos)
		args = append(args, *req.CompanyID)
		argPos++
	}

	search := req.Search
	if search == "" {
		search = req.Name
	}
	if search != "" {
		baseQuery += fmt.Sprintf(" AND (p.name ILIKE $%d OR p.description ILIKE $%d)", argPos, argPos)
		args = append(args, "%"+search+"%")
		argPos++
	}

	if req.Type != "" {
		baseQuery += fmt.Sprintf(" AND p.type = $%d", argPos)
		args = append(args, req.Type)
		argPos++
	}
	if req.MinPrice > 0 {
		baseQuery += fmt.Sprintf(" AND p.price >= $%d", argPos)
		args = append(args, req.MinPrice)
		argPos++
	}
	if req.MaxPrice > 0 {
		baseQuery += fmt.Sprintf(" AND p.price <= $%d", argPos)
		args = append(args, req.MaxPrice)
		argPos++
	}

	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	sortColumn := "p.created_at"
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

	rows, err := r.db.QueryContext(ctx, selectQuery+baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	products := make([]*models.Product, 0)
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.CompanyID, &p.CompanyName, &p.Name, &p.Description, &p.Type, &p.Price, &p.Stock, &p.ReservedStock, &p.ImagePath); err != nil {
			return nil, 0, err
		}
		products = append(products, &p)
	}

	return products, total, rows.Err()
}
