package dashboard

import (
	"encoding/json"
	"net/http"

	"crm-system-sales/internal/core/response"
	dashboardoverview "crm-system-sales/internal/services/dashboard_overview/service"
)

type DashboardController struct {
	service dashboardoverview.DashboardOverviewService
}

func NewDashboardController(service dashboardoverview.DashboardOverviewService) *DashboardController {
	return &DashboardController{service: service}
}

func (c *DashboardController) GetOverview(w http.ResponseWriter, r *http.Request) error {
	res, err := c.service.GetOverview(r.Context())
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}
