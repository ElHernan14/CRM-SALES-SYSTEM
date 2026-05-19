package invoiceaccess

import (
	constants "crm-system-sales/internal/modules/invoice/constants"

	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"

	clientModel "crm-system-sales/internal/models/client"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"

	"net/http"
)

func CanMutateDraftInvoice(
	tenant *tenantctx.TenantContext,
	invoice *invoiceModel.Invoice,
	buyer *clientModel.Client,
) error {

	if IsAdmin(tenant) {
		return nil
	}

	if invoice.StatusInvoice != constants.InvoiceDraft {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"Invoice no editable",
		)
	}

	// seller
	if tenant.CompanyID != nil &&
		*tenant.CompanyID == invoice.SellerCompanyID {

		return nil
	}

	// buyer company
	if buyer.CompanyID != nil &&
		tenant.CompanyID != nil &&
		*buyer.CompanyID == *tenant.CompanyID {

		return nil
	}

	// buyer individual
	if buyer.UserID != nil &&
		*buyer.UserID == tenant.UserID {

		return nil
	}

	return errorHandler.NewAppError(
		http.StatusForbidden,
		"No autorizado",
	)
}
