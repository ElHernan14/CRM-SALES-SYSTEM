package categories

import (
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterCategoriesRoutes(r *mux.Router, controller *CategoriesController) {
	r.HandleFunc("/categories", middleware.ErrorMiddleware(controller.GetAll)).Methods("GET")
	r.HandleFunc("/product-types", middleware.ErrorMiddleware(controller.GetProductTypes)).Methods("GET")
}
