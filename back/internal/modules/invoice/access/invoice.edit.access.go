package invoiceaccess

import (
	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"
	constants "crm-system-sales/internal/modules/invoice/constants"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"
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
