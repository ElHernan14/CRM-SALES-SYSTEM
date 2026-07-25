package marketplace

import (
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterMarketplaceRoutes(r *mux.Router, controller *MarketplaceController) {
	marketplace := r.PathPrefix("/marketplace").Subrouter()
	marketplace.HandleFunc("/suppliers", middleware.ErrorMiddleware(controller.GetSuppliers)).Methods("GET")
	marketplace.HandleFunc("/suppliers/{id}", middleware.ErrorMiddleware(controller.GetSupplierByID)).Methods("GET")
}
