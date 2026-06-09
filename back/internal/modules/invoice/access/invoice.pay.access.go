package invoiceaccess

import (
	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"
	constantInvoice "crm-system-sales/internal/modules/invoice/constants"

	clientModel "crm-system-sales/internal/models/client"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"

	"net/http"
)

func CanPayInvoice(
	tenant *tenantctx.TenantContext,
	invoice *invoiceModel.Invoice,
	buyer *clientModel.Client,
) error {
	if invoice.StatusInvoice != constantInvoice.InvoicePending {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"La invoice no está pendiente de pago",
		)
	}

	if IsAdmin(tenant) {
		return nil
	}

	// Buyer Company
	if buyer.CompanyID != nil &&
		tenant.CompanyID != nil &&
		*buyer.CompanyID == *tenant.CompanyID {

		return nil
	}

	// Buyer Individual

	if buyer.UserID != nil &&
		*buyer.UserID == tenant.UserID {

		return nil
	}

	return errorHandler.NewAppError(
		http.StatusForbidden,
		"No autorizado para pagar esta invoice",
	)
}
