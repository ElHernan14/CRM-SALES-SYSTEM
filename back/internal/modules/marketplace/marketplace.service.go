package marketplace

import (
	"context"
	coreDto "crm-system-sales/internal/core/dto"
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
	items, total, err := s.repo.ListSuppliers(ctx, req)
	if err != nil {
		return nil, err
	}

	return &marketplacedto.GetSuppliersResponse{
		Items: items,
		Meta: coreDto.Meta{
			Page:  req.Page,
			Limit: req.Limit,
			Total: total,
		},
	}, nil
}
