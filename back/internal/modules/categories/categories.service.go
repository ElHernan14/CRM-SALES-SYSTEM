package categories

import (
	"context"

	categoriesdto "crm-system-sales/internal/modules/categories/dto"
	categorycompany "crm-system-sales/internal/modules/category_company"
	categoryproduct "crm-system-sales/internal/modules/category_product"
)

type CategoriesService interface {
	GetAll(ctx context.Context) (*categoriesdto.GetCategoriesResponse, error)
}

type categoriesService struct {
	productRepo categoryproduct.CategoryProductRepository
	companyRepo categorycompany.CategoryCompanyRepository
}

func NewCategoriesService(
	productRepo categoryproduct.CategoryProductRepository,
	companyRepo categorycompany.CategoryCompanyRepository,
) CategoriesService {
	return &categoriesService{
		productRepo: productRepo,
		companyRepo: companyRepo,
	}
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

	res := &categoriesdto.GetCategoriesResponse{
		ProductCategories: make([]categoriesdto.CategoryItemResponse, 0, len(productCategories)),
		CompanyCategories: make([]categoriesdto.CategoryItemResponse, 0, len(companyCategories)),
	}

	for _, c := range productCategories {
		res.ProductCategories = append(res.ProductCategories, categoriesdto.CategoryItemResponse{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		})
	}

	for _, c := range companyCategories {
		res.CompanyCategories = append(res.CompanyCategories, categoriesdto.CategoryItemResponse{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		})
	}

	return res, nil
}
