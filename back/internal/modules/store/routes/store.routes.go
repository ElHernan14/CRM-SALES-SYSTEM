package routes

import (
	constants "crm-system-sales/internal/constants"
	"crm-system-sales/internal/middleware"
	controller "crm-system-sales/internal/modules/store/controller"

	"github.com/gorilla/mux"
)

func RegisterInvoiceItemRoutes(r *mux.Router, controller *controller.StoreController) {
	r.HandleFunc("/store/products", middleware.RequirePermission(constants.ProductRead)(middleware.ErrorMiddleware(controller.GetProducts))).Methods("GET")
	r.HandleFunc("/store/checkout", middleware.RequirePermission(constants.StoreCheckout)(middleware.ErrorMiddleware(controller.Checkout))).Methods("POST")
}
