package marketplace

import (
	"context"
	coreDto "crm-system-sales/internal/core/dto"
	tenantHelper "crm-system-sales/internal/core/tenant"
	marketplacedto "crm-system-sales/internal/modules/marketplace/dto"
)

type MarketplaceService interface {
	GetSuppliers(ctx context.Context, req *marketplacedto.GetSuppliersRequest) (*marketplacedto.GetSuppliersResponse, error)
}

type marketplaceService struct {
	repo MarketplaceRepository
}

func NewMarketplaceService(repo MarketplaceRepository) MarketplaceService {
	return &marketplaceService{repo: repo}
}

func (s *marketplaceService) GetSuppliers(ctx context.Context, req *marketplacedto.GetSuppliersRequest) (*marketplacedto.GetSuppliersResponse, error) {
	tenant := tenantHelper.GetTenant(ctx)
	var excludedCompanyID *int
	if tenant != nil {
		excludedCompanyID = tenant.CompanyID
	}

	items, total, err := s.repo.ListSuppliers(ctx, req, excludedCompanyID)
	if err != nil {
		return nil, err
	}

	return &marketplacedto.GetSuppliersResponse{
		Items: items,
		Meta:  coreDto.NewMeta(req.Page, req.Limit, total),
	}, nil
}
