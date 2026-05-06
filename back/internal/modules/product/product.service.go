package product

import (
	"context"
	"crm-system-sales/internal/core/access"
	meta "crm-system-sales/internal/core/dto"
	errorHandler "crm-system-sales/internal/core/error"
	tenant "crm-system-sales/internal/core/tenant"
	"crm-system-sales/internal/core/utils"
	models "crm-system-sales/internal/models/product"
	productdto "crm-system-sales/internal/modules/product/dto"
	"database/sql"
	"log"
	"net/http"
)

type ProductService interface {
	Create(ctx context.Context, req *productdto.CreateProductRequest) (*productdto.ProductResponse, error)
	GetProducts(ctx context.Context, req *productdto.GetProductsRequest) (*productdto.GetProductsResponse, error)
}

type productService struct {
	db   *sql.DB
	repo ProductRepository
}

func NewProductService(db *sql.DB, repo ProductRepository) ProductService {
	return &productService{db: db, repo: repo}
}

func (s *productService) Create(
	ctx context.Context,
	req *productdto.CreateProductRequest,
) (*productdto.ProductResponse, error) {

	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE CreateProduct")(err)
	}()

	tenant := tenant.GetTenant(ctx)

	if tenant == nil {
		return nil, errorHandler.NewAppError(http.StatusUnauthorized, "Usuario no autorizado")
	}

	companyID, err := access.ResolveCreateProductCompanyID(tenant, req.CompanyID)
	if err != nil {
		return nil, err
	}

	product := &models.Product{
		CompanyID:   *companyID,
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
		Price:       req.Price,
		Stock:       req.Stock,
		Status:      1,
	}

	err = s.repo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return &productdto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Type:        product.Type,
		Price:       product.Price,
		Stock:       product.Stock,
	}, nil
}

func (s *productService) GetProducts(
	ctx context.Context,
	req *productdto.GetProductsRequest,
) (*productdto.GetProductsResponse, error) {

	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE GetProducts")(err)
	}()

	tenant := tenant.GetTenant(ctx)
	if tenant == nil {
		log.Printf("Usuario no autorizado para ver productos.")
		return nil, errorHandler.NewAppError(http.StatusUnauthorized, "Usuario no autorizado para ver productos.")
	}

	offset := (req.Page - 1) * req.Limit

	companyID, err := access.ResolveGetProductsCompanyID(tenant, req.CompanyID)
	if err != nil {
		return nil, err
	}

	products, total, err := s.repo.GetAll(
		ctx,
		req.Search,
		req.Type,
		req.MinPrice,
		req.MaxPrice,
		companyID,
		req.Limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	var data []productdto.ProductListItem

	for _, p := range products {
		data = append(data, productdto.ProductListItem{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Type:        p.Type,
			Price:       p.Price,
			Stock:       p.Stock,
			Status:      p.Status,
			CompanyID:   p.CompanyID,
		})
	}

	return &productdto.GetProductsResponse{
		Data: data,
		Meta: meta.Meta{
			Page:  req.Page,
			Limit: req.Limit,
			Total: total,
		},
	}, nil
}
