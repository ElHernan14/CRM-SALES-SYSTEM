package invoiceaccess

import (
	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"
	invoiceModel "crm-system-sales/internal/modules/invoice/models"
	"net/http"
)

func CanManageInvoice(
	tenant *tenantctx.TenantContext,
	invoice *invoiceModel.Invoice,
) error {

	if IsAdmin(tenant) {
		return nil
	}

	if tenant.CompanyID == nil {
		return errorHandler.NewAppError(
			http.StatusForbidden,
			"No autorizado",
		)
	}

	if *tenant.CompanyID != invoice.SellerCompanyID {
		return errorHandler.NewAppError(
			http.StatusForbidden,
			"No autorizado para administrar esta invoice",
		)
	}

	return nil
}
