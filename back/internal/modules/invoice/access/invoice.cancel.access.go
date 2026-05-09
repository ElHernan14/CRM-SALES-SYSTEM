package invoiceaccess

import (
	constants "crm-system-sales/internal/constants"
	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"
	invoiceModel "crm-system-sales/internal/models/invoice"
	"net/http"
)

func CanCancelInvoice(
	tenant *tenantctx.TenantContext,
	invoice *invoiceModel.Invoice,
) error {

	if invoice.StatusInvoice == constants.InvoicePaid {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"No se puede cancelar una invoice pagada",
		)
	}

	return CanManageInvoice(tenant, invoice)
}
