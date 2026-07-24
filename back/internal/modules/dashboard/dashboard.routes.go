package dashboard

import (
	"crm-system-sales/internal/constants"
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterDashboardRoutes(r *mux.Router, controller *DashboardController) {
	r.HandleFunc("/dashboard/overview", middleware.RequirePermission(constants.CompanyRead)(middleware.ErrorMiddleware(controller.GetOverview))).Methods("GET")
}
