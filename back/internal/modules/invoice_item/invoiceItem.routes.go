package invoiceitem

import (
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterInvoiceItemRoutes(r *mux.Router, controller *InvoiceItemController) {
	r.HandleFunc("/invoice/{id}/items", middleware.RequirePermission("invoice_item:create")(middleware.ErrorMiddleware(controller.Create))).Methods("POST")
	r.HandleFunc("/invoice/{id}/items/{itemId}", middleware.RequirePermission("invoice_item:update")(middleware.ErrorMiddleware(controller.Update))).Methods("PATCH")
}
