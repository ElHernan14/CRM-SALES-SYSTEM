package categories

import (
	"context"

	categoriesdto "crm-system-sales/internal/modules/categories/dto"
	categorycompany "crm-system-sales/internal/modules/category_company"
	categoryproduct "crm-system-sales/internal/modules/category_product"
	producttype "crm-system-sales/internal/modules/product_type"
)

type CategoriesService interface {
	GetAll(ctx context.Context) (*categoriesdto.GetCategoriesResponse, error)
	GetProductTypes(ctx context.Context, categoryID *int) (*categoriesdto.GetProductTypesResponse, error)
}

type categoriesService struct {
	productRepo     categoryproduct.CategoryProductRepository
	companyRepo     categorycompany.CategoryCompanyRepository
	productTypeRepo producttype.ProductTypeRepository
}

func NewCategoriesService(
	productRepo categoryproduct.CategoryProductRepository,
	companyRepo categorycompany.CategoryCompanyRepository,
	productTypeRepo producttype.ProductTypeRepository,
) CategoriesService {
	return &categoriesService{productRepo: productRepo, companyRepo: companyRepo, productTypeRepo: productTypeRepo}
}

func (s *categoriesService) GetAll(ctx context.Context) (*categoriesdto.GetCategoriesResponse, error) {
	productCategories, err := s.productRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	companyCategories, err := s.companyRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	productTypes, err := s.productTypeRepo.GetAll(ctx, nil)
	if err != nil {
		return nil, err
	}

	res := &categoriesdto.GetCategoriesResponse{
		ProductCategories: make([]categoriesdto.CategoryItemResponse, 0, len(productCategories)),
		CompanyCategories: make([]categoriesdto.CategoryItemResponse, 0, len(companyCategories)),
		ProductTypes:      make([]categoriesdto.ProductTypeItemResponse, 0, len(productTypes)),
	}
	for _, c := range productCategories {
		res.ProductCategories = append(res.ProductCategories, categoriesdto.CategoryItemResponse{ID: c.ID, Name: c.Name, Description: c.Description})
	}
	for _, c := range companyCategories {
		res.CompanyCategories = append(res.CompanyCategories, categoriesdto.CategoryItemResponse{ID: c.ID, Name: c.Name, Description: c.Description})
	}
	for _, t := range productTypes {
		res.ProductTypes = append(res.ProductTypes, categoriesdto.ProductTypeItemResponse{ID: t.ID, CategoryID: t.CategoryID, Category: t.Category, Name: t.Name, Description: t.Description})
	}
	return res, nil
}

func (s *categoriesService) GetProductTypes(ctx context.Context, categoryID *int) (*categoriesdto.GetProductTypesResponse, error) {
	productTypes, err := s.productTypeRepo.GetAll(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	res := &categoriesdto.GetProductTypesResponse{Items: make([]categoriesdto.ProductTypeItemResponse, 0, len(productTypes))}
	for _, t := range productTypes {
		res.Items = append(res.Items, categoriesdto.ProductTypeItemResponse{ID: t.ID, CategoryID: t.CategoryID, Category: t.Category, Name: t.Name, Description: t.Description})
	}
	return res, nil
}
