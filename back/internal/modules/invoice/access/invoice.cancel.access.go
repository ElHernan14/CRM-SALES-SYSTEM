package invoiceaccess

import (
	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"
	constants "crm-system-sales/internal/modules/invoice/constants"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"
	"net/http"
)

func CanCancelInvoice(
	tenant *tenantctx.TenantContext,
	invoice *invoiceModel.Invoice,
) error {

	if IsAdmin(tenant) {
		return nil
	}

	if invoice.StatusInvoice == constants.InvoicePaid {

		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"No se puede cancelar una invoice pagada",
		)
	}

	if invoice.StatusInvoice == constants.InvoiceCanceled {

		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"La invoice ya está cancelada",
		)
	}

	if tenant.CompanyID != nil &&
		*tenant.CompanyID == invoice.SellerCompanyID {

		return nil
	}

	return errorHandler.NewAppError(
		http.StatusForbidden,
		"No autorizado",
	)
}
