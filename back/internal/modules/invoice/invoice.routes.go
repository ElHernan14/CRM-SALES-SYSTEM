package invoice

import (
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterInvoiceRoutes(r *mux.Router, controller *InvoiceController) {
	r.HandleFunc("/invoice", middleware.RequirePermission("invoice:create")(middleware.ErrorMiddleware(controller.CreateDraft))).Methods("POST")
	r.HandleFunc("/invoices/{id}/submit", middleware.RequirePermission("invoice:update")(middleware.ErrorMiddleware(controller.Submit))).Methods("POST")
}
