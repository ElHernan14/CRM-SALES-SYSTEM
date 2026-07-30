package product

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	models "crm-system-sales/internal/models/product"
	productdto "crm-system-sales/internal/modules/product/dto"
	storedto "crm-system-sales/internal/modules/store/dto"
)

type ProductRepository interface {
	Create(ctx context.Context, p *models.Product) error
	GetAll(ctx context.Context, search string, kind string, categoryID *int, category string, typeID *int, productType string, status *int, minPrice float64, maxPrice float64, companyID *int, limit int, offset int, sortColumn string, order string) ([]models.Product, int, error)
	GetByID(ctx context.Context, id int) (*models.Product, error)
	Update(ctx context.Context, p *models.Product) error
	UpdateImage(ctx context.Context, productID int, imagePath string) error
	SoftDelete(ctx context.Context, id int) error
	GetByIDForUpdate(ctx context.Context, tx *sql.Tx, id int) (*models.Product, error)
	ListByCompanyID(ctx context.Context, companyID int, req *productdto.GetCompanyProductsRequest) ([]*models.Product, int, error)
	ListAvailableProducts(ctx context.Context, req *storedto.GetStoreProductsRequest) ([]*models.Product, int, error)
	GetAvailableProductByID(ctx context.Context, id int) (*models.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, p *models.Product) error {
	query := `
		INSERT INTO product (company_id, name, description, kind, category_id, category, type_id, type, price, stock, image_path, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, 1)
		RETURNING id
	`

	return r.db.QueryRowContext(ctx, query, p.CompanyID, p.Name, p.Description, p.Kind, p.CategoryID, p.Category, p.TypeID, p.Type, p.Price, p.Stock, p.ImagePath).Scan(&p.ID)
}

func (r *productRepository) GetAll(ctx context.Context, search string, kind string, categoryID *int, category string, typeID *int, productType string, status *int, minPrice float64, maxPrice float64, companyID *int, limit int, offset int, sortColumn string, order string) ([]models.Product, int, error) {
	baseQuery := `
		FROM product p
		INNER JOIN category_product cat ON cat.id = p.category_id
		INNER JOIN product_type pt ON pt.id = p.type_id
		WHERE true = true
	`

	args := []interface{}{}
	i := 1

	if companyID != nil {
		baseQuery += fmt.Sprintf(" AND p.company_id = $%d", i)
		args = append(args, *companyID)
		i++
	}
	if search != "" {
		baseQuery += fmt.Sprintf(` AND LOWER(p.name) LIKE LOWER($%d)`, i)
		args = append(args, "%"+search+"%")
		i++
	}
	if kind != "" {
		baseQuery += fmt.Sprintf(" AND p.kind = $%d", i)
		args = append(args, kind)
		i++
	}
	if categoryID != nil {
		baseQuery += fmt.Sprintf(" AND p.category_id = $%d", i)
		args = append(args, *categoryID)
		i++
	} else if category != "" {
		baseQuery += fmt.Sprintf(" AND LOWER(cat.name) = LOWER($%d)", i)
		args = append(args, category)
		i++
	}
	if typeID != nil {
		baseQuery += fmt.Sprintf(" AND p.type_id = $%d", i)
		args = append(args, *typeID)
		i++
	} else if productType != "" {
		baseQuery += fmt.Sprintf(" AND LOWER(pt.name) = LOWER($%d)", i)
		args = append(args, productType)
		i++
	}
	if status != nil {
		baseQuery += fmt.Sprintf(" AND p.status = $%d", i)
		args = append(args, *status)
		i++
		if *status == 1 {
			baseQuery += " AND p.deleted_at IS NULL"
		}
	}
	if minPrice > 0 {
		baseQuery += fmt.Sprintf(" AND p.price >= $%d", i)
		args = append(args, minPrice)
		i++
	}
	if maxPrice > 0 {
		baseQuery += fmt.Sprintf(" AND p.price <= $%d", i)
		args = append(args, maxPrice)
		i++
	}

	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	allowedSortColumns := map[string]string{
		"id":              "p.id",
		"name":            "p.name",
		"description":     "p.description",
		"kind":            "p.kind",
		"type":            "pt.name",
		"category":        "cat.name",
		"price":           "p.price",
		"stock":           "p.stock",
		"reserved_stock":  "p.reserved_stock",
		"available_stock": "(p.stock - p.reserved_stock)",
		"status":          "p.status",
		"created_at":      "p.created_at",
	}
	allowedOrders := map[string]string{"asc": "ASC", "desc": "DESC"}

	sortSQL := "p.id"
	orderSQL := "DESC"
	if v, ok := allowedSortColumns[sortColumn]; ok {
		sortSQL = v
	}
	if v, ok := allowedOrders[strings.ToLower(order)]; ok {
		orderSQL = v
	}

	dataQuery := `
		SELECT p.id, p.name, p.description, p.kind, p.type_id, pt.name, p.category_id, cat.name, p.price, p.stock, p.status, p.company_id, p.reserved_stock, p.image_path
	` + baseQuery + fmt.Sprintf(`
		ORDER BY %s %s, p.id DESC
		LIMIT $%d OFFSET $%d
	`, sortSQL, orderSQL, i, i+1)

	args = append(args, limit, offset)
	rows, err := r.db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Kind, &p.TypeID, &p.Type, &p.CategoryID, &p.Category, &p.Price, &p.Stock, &p.Status, &p.CompanyID, &p.ReservedStock, &p.ImagePath); err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	return products, total, rows.Err()
}

func (r *productRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	query := `
		SELECT p.id, p.name, p.description, p.kind, p.type_id, pt.name, p.category_id, cat.name, p.price, p.stock, p.reserved_stock, p.status, p.company_id, p.image_path
		FROM product p
		INNER JOIN category_product cat ON cat.id = p.category_id
		INNER JOIN product_type pt ON pt.id = p.type_id
		WHERE p.id = $1
	`

	var p models.Product
	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Kind, &p.TypeID, &p.Type, &p.CategoryID, &p.Category, &p.Price, &p.Stock, &p.ReservedStock, &p.Status, &p.CompanyID, &p.ImagePath)
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
			kind = $3,
			category_id = $4,
			category = $5,
			type_id = $6,
			type = $7,
			price = $8,
			stock = $9,
			status = $10
		WHERE id = $11
	`

	_, err := r.db.ExecContext(ctx, query, p.Name, p.Description, p.Kind, p.CategoryID, p.Category, p.TypeID, p.Type, p.Price, p.Stock, p.Status, p.ID)
	return err
}

func (r *productRepository) UpdateImage(ctx context.Context, productID int, imagePath string) error {
	query := `UPDATE product SET image_path = $1 WHERE id = $2 `
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
	_, err := r.db.ExecContext(ctx, `UPDATE product SET deleted_at = NOW(), status = 0 WHERE id = $1`, id)
	return err
}

func (r *productRepository) GetByIDForUpdate(ctx context.Context, tx *sql.Tx, id int) (*models.Product, error) {
	query := `SELECT id, name, price, stock, reserved_stock, company_id FROM product WHERE id = $1 FOR UPDATE`
	product := &models.Product{}
	err := tx.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.ReservedStock, &product.CompanyID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) ListByCompanyID(ctx context.Context, companyID int, req *productdto.GetCompanyProductsRequest) ([]*models.Product, int, error) {
	allowedSortColumns := map[string]string{"name": "p.name", "description": "p.description", "kind": "p.kind", "type": "pt.name", "category": "cat.name", "price": "p.price", "stock": "p.stock"}
	allowedOrders := map[string]string{"asc": "ASC", "desc": "DESC"}

	selectQuery := `SELECT p.id, p.name, p.description, p.kind, p.type_id, pt.name, p.category_id, cat.name, p.price, p.stock, p.reserved_stock, p.status, p.image_path `
	baseQuery := `
		FROM product p
		INNER JOIN category_product cat ON cat.id = p.category_id
		INNER JOIN product_type pt ON pt.id = p.type_id
		WHERE p.company_id = $1
	`
	args := []interface{}{companyID}
	argPos := 2

	if req.Name != "" {
		baseQuery += fmt.Sprintf(" AND p.name ILIKE $%d", argPos)
		args = append(args, "%"+req.Name+"%")
		argPos++
	}
	if req.Kind != "" {
		baseQuery += fmt.Sprintf(" AND p.kind = $%d", argPos)
		args = append(args, req.Kind)
		argPos++
	}
	if req.CategoryID != nil {
		baseQuery += fmt.Sprintf(" AND p.category_id = $%d", argPos)
		args = append(args, *req.CategoryID)
		argPos++
	} else if req.Category != "" {
		baseQuery += fmt.Sprintf(" AND LOWER(cat.name) = LOWER($%d)", argPos)
		args = append(args, req.Category)
		argPos++
	}
	if req.TypeID != nil {
		baseQuery += fmt.Sprintf(" AND p.type_id = $%d", argPos)
		args = append(args, *req.TypeID)
		argPos++
	} else if req.Type != "" {
		baseQuery += fmt.Sprintf(" AND LOWER(pt.name) = LOWER($%d)", argPos)
		args = append(args, req.Type)
		argPos++
	}
	if req.Status != nil {
		baseQuery += fmt.Sprintf(" AND p.status = $%d", argPos)
		args = append(args, *req.Status)
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
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Kind, &p.TypeID, &p.Type, &p.CategoryID, &p.Category, &p.Price, &p.Stock, &p.ReservedStock, &p.Status, &p.ImagePath); err != nil {
			return nil, 0, err
		}
		products = append(products, &p)
	}
	return products, total, rows.Err()
}

func (r *productRepository) ListAvailableProducts(ctx context.Context, req *storedto.GetStoreProductsRequest) ([]*models.Product, int, error) {
	allowedSortColumns := map[string]string{"name": "p.name", "created_at": "p.created_at", "kind": "p.kind", "type": "pt.name", "category": "cat.name", "price": "p.price", "stock": "p.stock", "reserved_stock": "p.reserved_stock", "available_stock": "(p.stock - p.reserved_stock)"}
	allowedOrders := map[string]string{"asc": "ASC", "desc": "DESC"}

	selectQuery := `
		SELECT p.id, p.company_id, c.name AS company_name, p.name, p.description, p.kind, p.type_id, pt.name, p.category_id, cat.name, p.price, p.stock, p.reserved_stock, p.image_path
	`
	baseQuery := `
		FROM product p
		INNER JOIN company c ON c.id = p.company_id
		INNER JOIN category_product cat ON cat.id = p.category_id
		INNER JOIN product_type pt ON pt.id = p.type_id
		WHERE p.status = 1
		  AND p.deleted_at IS NULL
		  AND c.status = 1
		  AND c.deleted_at IS NULL
	`

	args := []interface{}{}
	argPos := 1
	if req.CompanyID != nil {
		baseQuery += fmt.Sprintf(" AND p.company_id = $%d", argPos)
		args = append(args, *req.CompanyID)
		argPos++
	}
	if req.ExcludedCompanyID != nil {
		baseQuery += fmt.Sprintf(" AND p.company_id <> $%d", argPos)
		args = append(args, *req.ExcludedCompanyID)
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
	if req.Kind != "" {
		baseQuery += fmt.Sprintf(" AND p.kind = $%d", argPos)
		args = append(args, req.Kind)
		argPos++
	}
	if req.CategoryID != nil {
		baseQuery += fmt.Sprintf(" AND p.category_id = $%d", argPos)
		args = append(args, *req.CategoryID)
		argPos++
	} else if req.Category != "" {
		baseQuery += fmt.Sprintf(" AND LOWER(cat.name) = LOWER($%d)", argPos)
		args = append(args, req.Category)
		argPos++
	}
	if req.TypeID != nil {
		baseQuery += fmt.Sprintf(" AND p.type_id = $%d", argPos)
		args = append(args, *req.TypeID)
		argPos++
	} else if req.Type != "" {
		baseQuery += fmt.Sprintf(" AND LOWER(pt.name) = LOWER($%d)", argPos)
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
		if err := rows.Scan(&p.ID, &p.CompanyID, &p.CompanyName, &p.Name, &p.Description, &p.Kind, &p.TypeID, &p.Type, &p.CategoryID, &p.Category, &p.Price, &p.Stock, &p.ReservedStock, &p.ImagePath); err != nil {
			return nil, 0, err
		}
		products = append(products, &p)
	}
	return products, total, rows.Err()
}

func (r *productRepository) GetAvailableProductByID(ctx context.Context, id int) (*models.Product, error) {
	query := `
		SELECT
			p.id,
			p.company_id,
			c.name AS company_name,
			p.name,
			p.description,
			p.kind,
			p.type_id,
			pt.name AS type_name,
			p.category_id,
			cat.name AS category_name,
			p.price,
			p.stock,
			p.reserved_stock,
			p.image_path
		FROM product p
		INNER JOIN company c ON c.id = p.company_id
		INNER JOIN category_product cat ON cat.id = p.category_id
		INNER JOIN product_type pt ON pt.id = p.type_id
		WHERE p.id = $1
		  AND p.status = 1
		  AND p.deleted_at IS NULL
		  AND c.status = 1
		  AND c.deleted_at IS NULL
	`

	var p models.Product
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.CompanyID,
		&p.CompanyName,
		&p.Name,
		&p.Description,
		&p.Kind,
		&p.TypeID,
		&p.Type,
		&p.CategoryID,
		&p.Category,
		&p.Price,
		&p.Stock,
		&p.ReservedStock,
		&p.ImagePath,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}
