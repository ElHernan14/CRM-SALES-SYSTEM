package invoiceaccess

import (
	constants "crm-system-sales/internal/constants"
	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"
	invoiceModel "crm-system-sales/internal/models/invoice"
	"net/http"
)

func CanEditDraftInvoice(
	tenant *tenantctx.TenantContext,
	invoice *invoiceModel.Invoice,
) error {

	if IsAdmin(tenant) {
		return nil
	}

	if invoice.StatusInvoice != constants.InvoiceDraft {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"La invoice no está en draft",
		)
	}

	if invoice.CreatedByUserID != tenant.UserID {
		return errorHandler.NewAppError(
			http.StatusForbidden,
			"No autorizado para editar esta invoice",
		)
	}

	return nil
}
