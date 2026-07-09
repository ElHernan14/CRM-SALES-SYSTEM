package marketplace

import (
	"crm-system-sales/internal/constants"
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterMarketplaceRoutes(r *mux.Router, controller *MarketplaceController) {
	marketplace := r.PathPrefix("/marketplace").Subrouter()
	marketplace.HandleFunc("/suppliers", middleware.RequirePermission(constants.ProductRead)(middleware.ErrorMiddleware(controller.GetSuppliers))).Methods("GET")
}
