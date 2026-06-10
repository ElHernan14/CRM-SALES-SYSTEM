package invoicepayment

import (
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterInvoicePaymentRoutes(r *mux.Router, controller *InvoicePaymentController) {
	r.HandleFunc("/invoice/{id}/payments", middleware.RequirePermission("invoice_payment:read")(middleware.ErrorMiddleware(controller.GetByInvoice))).Methods("GET")
}
