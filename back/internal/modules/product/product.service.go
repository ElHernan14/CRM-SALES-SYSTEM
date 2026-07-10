package product

import (
	"context"
	"database/sql"
	"mime/multipart"
	"net/http"

	"crm-system-sales/internal/core/access"
	"crm-system-sales/internal/core/dto"
	meta "crm-system-sales/internal/core/dto"
	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/files"
	tenant "crm-system-sales/internal/core/tenant"
	models "crm-system-sales/internal/models/product"
	productdto "crm-system-sales/internal/modules/product/dto"
)

type ProductService interface {
	Create(ctx context.Context, req *productdto.CreateProductRequest) (*productdto.ProductResponse, error)
	GetProducts(ctx context.Context, req *productdto.GetProductsRequest) (*productdto.GetProductsResponse, error)
	GetByID(ctx context.Context, id int) (*productdto.ProductDetailResponse, error)
	Update(ctx context.Context, id int, req *productdto.UpdateProductRequest) (*productdto.ProductDetailResponse, error)
	UploadImage(ctx context.Context, id int, file multipart.File, header *multipart.FileHeader) (*productdto.UploadProductImageResponse, error)
	Delete(ctx context.Context, id int) error
	GetCompanyProducts(ctx context.Context, req *productdto.GetCompanyProductsRequest) (*productdto.GetCompanyProductsResponse, error)
}

type productService struct {
	db           *sql.DB
	repo         ProductRepository
	imageStorage files.ImageStorage
}

func NewProductService(db *sql.DB, repo ProductRepository, imageStorage files.ImageStorage) ProductService {
	return &productService{db: db, repo: repo, imageStorage: imageStorage}
}

func (s *productService) Create(ctx context.Context, req *productdto.CreateProductRequest) (*productdto.ProductResponse, error) {
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
		ImagePath:   req.ImagePath,
		Status:      1,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return mapProductResponse(product), nil
}

func (s *productService) GetProducts(ctx context.Context, req *productdto.GetProductsRequest) (*productdto.GetProductsResponse, error) {
	tenant := tenant.GetTenant(ctx)
	if tenant == nil {
		return nil, errorHandler.NewAppError(http.StatusUnauthorized, "Usuario no autorizado para ver productos.")
	}

	offset := (req.Page - 1) * req.Limit
	companyID, err := access.ResolveGetProductsCompanyID(tenant, req.CompanyID)
	if err != nil {
		return nil, err
	}

	products, total, err := s.repo.GetAll(ctx, req.Search, req.Type, req.MinPrice, req.MaxPrice, companyID, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	data := make([]productdto.ProductListItem, 0, len(products))
	for _, p := range products {
		available := p.Stock - p.ReservedStock
		data = append(data, productdto.ProductListItem{
			ID:             p.ID,
			Name:           p.Name,
			Description:    p.Description,
			Type:           p.Type,
			Price:          p.Price,
			Stock:          p.Stock,
			Status:         p.Status,
			CompanyID:      p.CompanyID,
			AvailableStock: available,
			ImagePath:      p.ImagePath,
		})
	}

	return &productdto.GetProductsResponse{Data: data, Meta: meta.NewMeta(req.Page, req.Limit, total)}, nil
}

func (s *productService) GetByID(ctx context.Context, id int) (*productdto.ProductDetailResponse, error) {
	product, err := s.getOwnedProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapProductDetailResponse(product), nil
}

func (s *productService) Update(ctx context.Context, id int, req *productdto.UpdateProductRequest) (*productdto.ProductDetailResponse, error) {
	product, err := s.getOwnedProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.Type != nil {
		product.Type = *req.Type
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.Stock != nil {
		product.Stock = *req.Stock
	}
	if req.Status != nil {
		product.Status = *req.Status
	}

	if err := s.repo.Update(ctx, product); err != nil {
		return nil, err
	}

	return mapProductDetailResponse(product), nil
}

func (s *productService) UploadImage(ctx context.Context, id int, file multipart.File, header *multipart.FileHeader) (*productdto.UploadProductImageResponse, error) {
	product, err := s.getOwnedProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	imagePath, err := s.imageStorage.SaveImage(file, header, "product/images", "product", product.ID)
	if err != nil {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, err.Error())
	}

	if err := s.repo.UpdateImage(ctx, product.ID, imagePath); err != nil {
		if err == sql.ErrNoRows {
			return nil, errorHandler.NewAppError(http.StatusNotFound, "Producto no encontrado")
		}
		return nil, err
	}

	return &productdto.UploadProductImageResponse{ImagePath: imagePath}, nil
}

func (s *productService) Delete(ctx context.Context, id int) error {
	product, err := s.getOwnedProduct(ctx, id)
	if err != nil {
		return err
	}

	if product.Status == 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "Producto ya fue eliminado")
	}

	return s.repo.SoftDelete(ctx, id)
}

func (s *productService) GetCompanyProducts(ctx context.Context, req *productdto.GetCompanyProductsRequest) (*productdto.GetCompanyProductsResponse, error) {
	tenant := tenant.GetTenant(ctx)
	if tenant == nil || tenant.CompanyID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "Solo empresas")
	}

	products, total, err := s.repo.ListByCompanyID(ctx, *tenant.CompanyID, req)
	if err != nil {
		return nil, err
	}

	items := make([]productdto.CompanyProductResponse, 0, len(products))
	for _, p := range products {
		items = append(items, productdto.CompanyProductResponse{
			ID:            p.ID,
			Name:          p.Name,
			Description:   p.Description,
			Type:          p.Type,
			Price:         p.Price,
			Stock:         p.Stock,
			ReservedStock: p.ReservedStock,
			Status:        p.Status,
			ImagePath:     p.ImagePath,
		})
	}

	return &productdto.GetCompanyProductsResponse{Items: items, Meta: dto.NewMeta(req.Page, req.Limit, total)}, nil
}

func (s *productService) getOwnedProduct(ctx context.Context, id int) (*models.Product, error) {
	tenant := tenant.GetTenant(ctx)
	if tenant == nil {
		return nil, errorHandler.NewAppError(http.StatusUnauthorized, "Usuario no autorizado")
	}

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errorHandler.NewAppError(http.StatusNotFound, "Producto no encontrado")
	}

	if _, err := access.ResolveGetProductsCompanyID(tenant, &product.CompanyID); err != nil {
		return nil, err
	}

	return product, nil
}

func mapProductResponse(product *models.Product) *productdto.ProductResponse {
	available := product.Stock - product.ReservedStock
	return &productdto.ProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		Description:    product.Description,
		Type:           product.Type,
		Price:          product.Price,
		Stock:          product.Stock,
		ReservedStock:  product.ReservedStock,
		AvailableStock: available,
		ImagePath:      product.ImagePath,
	}
}

func mapProductDetailResponse(product *models.Product) *productdto.ProductDetailResponse {
	return &productdto.ProductDetailResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Type:        product.Type,
		Price:       product.Price,
		Stock:       product.Stock,
		Status:      product.Status,
		CompanyID:   product.CompanyID,
		ImagePath:   product.ImagePath,
	}
}
