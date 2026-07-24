package service

import (
	"context"
	"net/http"

	errorHandler "crm-system-sales/internal/core/error"
	tenanthelper "crm-system-sales/internal/core/tenant"
	dashboarddto "crm-system-sales/internal/modules/dashboard/dto"
)

type DashboardOverviewService interface {
	GetOverview(ctx context.Context) (*dashboarddto.DashboardOverviewResponse, error)
}

type DashboardOverviewRepository interface {
	GetSalesSummary(ctx context.Context, companyID int) (dashboarddto.DashboardSalesSummaryResponse, error)
	GetPurchasesSummary(ctx context.Context, companyID int) (dashboarddto.DashboardPurchasesSummaryResponse, error)
	GetInventorySummary(ctx context.Context, companyID int) (dashboarddto.DashboardInventorySummaryResponse, error)
	CountPartiallyPaidPurchases(ctx context.Context, companyID int) (int, error)
	GetRecentSales(ctx context.Context, companyID int, limit int) ([]dashboarddto.DashboardRecentInvoiceResponse, error)
	GetRecentPurchases(ctx context.Context, companyID int, limit int) ([]dashboarddto.DashboardRecentInvoiceResponse, error)
	GetLowStockProducts(ctx context.Context, companyID int, limit int) ([]dashboarddto.DashboardLowStockProductResponse, error)
}

type dashboardOverviewService struct {
	repo DashboardOverviewRepository
}

func NewDashboardOverviewService(repo DashboardOverviewRepository) DashboardOverviewService {
	return &dashboardOverviewService{repo: repo}
}

func (s *dashboardOverviewService) GetOverview(ctx context.Context) (*dashboarddto.DashboardOverviewResponse, error) {
	tenant := tenanthelper.GetTenant(ctx)
	if tenant == nil || tenant.CompanyID == nil {
		return nil, errorHandler.NewAppError(http.StatusForbidden, "usuario no pertenece a una empresa")
	}

	companyID := *tenant.CompanyID

	sales, err := s.repo.GetSalesSummary(ctx, companyID)
	if err != nil {
		return nil, err
	}

	purchases, err := s.repo.GetPurchasesSummary(ctx, companyID)
	if err != nil {
		return nil, err
	}

	inventory, err := s.repo.GetInventorySummary(ctx, companyID)
	if err != nil {
		return nil, err
	}

	recentSales, err := s.repo.GetRecentSales(ctx, companyID, 5)
	if err != nil {
		return nil, err
	}

	recentPurchases, err := s.repo.GetRecentPurchases(ctx, companyID, 5)
	if err != nil {
		return nil, err
	}

	lowStockProducts, err := s.repo.GetLowStockProducts(ctx, companyID, 5)
	if err != nil {
		return nil, err
	}

	partiallyPaidPurchases, err := s.repo.CountPartiallyPaidPurchases(ctx, companyID)
	if err != nil {
		return nil, err
	}

	attention := dashboarddto.DashboardAttentionResponse{
		PendingSales:           sales.Pending,
		DraftSales:             sales.Draft,
		PendingPurchases:       purchases.Pending,
		PartiallyPaidPurchases: partiallyPaidPurchases,
		LowStockProducts:       inventory.LowStockProducts,
		OutOfStockProducts:     inventory.OutOfStockProducts,
	}

	return &dashboarddto.DashboardOverviewResponse{
		Sales:            sales,
		Purchases:        purchases,
		Inventory:        inventory,
		Attention:        attention,
		RecentSales:      recentSales,
		RecentPurchases:  recentPurchases,
		LowStockProducts: lowStockProducts,
	}, nil
}
