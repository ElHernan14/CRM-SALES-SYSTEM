package marketplace

import (
	"context"
	coreDto "crm-system-sales/internal/core/dto"
	errorHandler "crm-system-sales/internal/core/error"
	tenantHelper "crm-system-sales/internal/core/tenant"
	marketplacedto "crm-system-sales/internal/modules/marketplace/dto"
	"net/http"
)

type MarketplaceService interface {
	GetSuppliers(ctx context.Context, req *marketplacedto.GetSuppliersRequest) (*marketplacedto.GetSuppliersResponse, error)
	GetSupplierByID(ctx context.Context, supplierID int) (*marketplacedto.SupplierResponse, error)
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

func (s *marketplaceService) GetSupplierByID(ctx context.Context, supplierID int) (*marketplacedto.SupplierResponse, error) {
	if supplierID <= 0 {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, "supplier_id invalido")
	}

	tenant := tenantHelper.GetTenant(ctx)
	var excludedCompanyID *int
	if tenant != nil {
		excludedCompanyID = tenant.CompanyID
	}

	supplier, err := s.repo.GetSupplierByID(ctx, supplierID, excludedCompanyID)
	if err != nil {
		return nil, err
	}
	if supplier == nil {
		return nil, errorHandler.NewAppError(http.StatusNotFound, "supplier no encontrado")
	}

	return supplier, nil
}
