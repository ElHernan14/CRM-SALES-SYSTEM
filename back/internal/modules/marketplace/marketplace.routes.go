package marketplace

import (
	"crm-system-sales/internal/constants"
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterMarketplaceRoutes(r *mux.Router, controller *MarketplaceController) {
	marketplace := r.PathPrefix("/marketplace").Subrouter()
	marketplace.Handle("/suppliers", middleware.OptionalAuthMiddleware(middleware.ErrorMiddleware(controller.GetSuppliers))).Methods("GET")
	marketplace.Handle("/suppliers/{id}", middleware.OptionalAuthMiddleware(middleware.ErrorMiddleware(controller.GetSupplierByID))).Methods("GET")
}

func RegisterMarketplaceProtectedRoutes(r *mux.Router, controller *MarketplaceController) {
	marketplace := r.PathPrefix("/marketplace").Subrouter()
	marketplace.HandleFunc("/cart", middleware.RequirePermission(constants.InvoiceRead)(middleware.ErrorMiddleware(controller.GetCart))).Methods("GET")
	marketplace.HandleFunc("/cart/ensure", middleware.RequirePermission(constants.InvoiceCreate)(middleware.ErrorMiddleware(controller.EnsureCart))).Methods("POST")
	marketplace.HandleFunc("/checkout", middleware.RequirePermission(constants.InvoiceUpdate)(middleware.ErrorMiddleware(controller.Checkout))).Methods("POST")
}
