package product

import (
	"crm-system-sales/internal/constants"
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router, controller *ProductController) {
	products := r.PathPrefix("/product").Subrouter()
	products.HandleFunc("/products", middleware.RequirePermission(constants.ProductCreate)(middleware.ErrorMiddleware(controller.CreateProduct))).Methods("POST")
	products.HandleFunc("/products", middleware.RequirePermission(constants.ProductView)(middleware.ErrorMiddleware(controller.GetProducts))).Methods("GET")
}
