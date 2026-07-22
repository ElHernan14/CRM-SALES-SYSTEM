package routes

import (
	constants "crm-system-sales/internal/constants"
	"crm-system-sales/internal/middleware"
	controller "crm-system-sales/internal/modules/store/controller"

	"github.com/gorilla/mux"
)

func RegisterInvoiceItemRoutes(r *mux.Router, controller *controller.StoreController) {
	r.HandleFunc("/store/products", middleware.RequirePermission(constants.ProductRead)(middleware.ErrorMiddleware(controller.GetProducts))).Methods("GET")
	r.HandleFunc("/store/products/{id}", middleware.RequirePermission(constants.ProductRead)(middleware.ErrorMiddleware(controller.GetProductByID))).Methods("GET")
	r.HandleFunc("/store/cart", middleware.RequirePermission(constants.ProductRead)(middleware.ErrorMiddleware(controller.GetCart))).Methods("GET")
	r.HandleFunc("/store/carts", middleware.RequirePermission(constants.ProductRead)(middleware.ErrorMiddleware(controller.GetCarts))).Methods("GET")
	r.HandleFunc("/store/cart/ensure", middleware.RequirePermission(constants.ProductRead)(middleware.ErrorMiddleware(controller.EnsureCart))).Methods("POST")
	r.HandleFunc("/store/checkout/all", middleware.RequirePermission(constants.StoreCheckout)(middleware.ErrorMiddleware(controller.CheckoutAll))).Methods("POST")
	r.HandleFunc("/store/checkout", middleware.RequirePermission(constants.StoreCheckout)(middleware.ErrorMiddleware(controller.Checkout))).Methods("POST")
}
