package invoiceRoutes

import (
	"crm-system-sales/internal/middleware"
	controller "crm-system-sales/internal/modules/invoice/controller"

	"github.com/gorilla/mux"
)

func RegisterInvoiceRoutes(r *mux.Router, controller *controller.InvoiceController) {
	r.HandleFunc("/invoice", middleware.RequirePermission("invoice:create")(middleware.ErrorMiddleware(controller.CreateDraft))).Methods("POST")
	r.HandleFunc("/invoices/{id}/submit", middleware.RequirePermission("invoice:update")(middleware.ErrorMiddleware(controller.Submit))).Methods("POST")
}
