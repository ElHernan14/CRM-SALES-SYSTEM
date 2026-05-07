package product

import (
	"crm-system-sales/internal/constants"
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router, controller *ProductController) {
	r.HandleFunc("/product", middleware.RequirePermission(constants.ProductCreate)(middleware.ErrorMiddleware(controller.CreateProduct))).Methods("POST")
	r.HandleFunc("/product", middleware.RequirePermission(constants.ProductView)(middleware.ErrorMiddleware(controller.GetProducts))).Methods("GET")
	r.HandleFunc("/product/{id}", middleware.RequirePermission(constants.ProductRead)(middleware.ErrorMiddleware(controller.GetProductByID))).Methods("GET")
	r.HandleFunc("/product/{id}", middleware.RequirePermission(constants.ProductUpdate)(middleware.ErrorMiddleware(controller.UpdateProduct))).Methods("PATCH")
	r.HandleFunc("/product/{id}", middleware.RequirePermission(constants.ProductDelete)(middleware.ErrorMiddleware(controller.DeleteProduct))).Methods("DELETE")
}
