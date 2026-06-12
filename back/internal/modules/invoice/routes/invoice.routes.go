package invoiceRoutes

import (
	"crm-system-sales/internal/middleware"
	controller "crm-system-sales/internal/modules/invoice/controller"

	"github.com/gorilla/mux"
)

func RegisterInvoiceRoutes(r *mux.Router, controller *controller.InvoiceController) {
	// State Machine
	r.HandleFunc("/invoice", middleware.RequirePermission("invoice:create")(middleware.ErrorMiddleware(controller.CreateDraft))).Methods("POST")
	r.HandleFunc("/invoice/{id}/submit", middleware.RequirePermission("invoice:submit")(middleware.ErrorMiddleware(controller.Submit))).Methods("POST")
	r.HandleFunc("/invoice/{id}/pay", middleware.RequirePermission("invoice:pay")(middleware.ErrorMiddleware(controller.Pay))).Methods("POST")
	r.HandleFunc("/invoice/{id}", middleware.RequirePermission("invoice:read")(middleware.ErrorMiddleware(controller.GetByID))).Methods("GET")
	r.HandleFunc("/invoices/{id}/cancel", middleware.RequirePermission("invoice:cancel")(middleware.ErrorMiddleware(controller.Cancel))).Methods("POST")

	// ERP Machine
	r.HandleFunc("/company/invoices", middleware.RequirePermission("invoice:read")(middleware.ErrorMiddleware(controller.GetCompanyInvoices))).Methods("GET")
}
