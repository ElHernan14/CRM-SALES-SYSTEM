package invoiceaccess

import (
	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"
	clientModel "crm-system-sales/internal/models/client"
	invoiceModel "crm-system-sales/internal/models/invoice"
	"net/http"
)

func CanViewInvoice(
	tenant *tenantctx.TenantContext,
	invoice *invoiceModel.Invoice,
	buyer *clientModel.Client,
) error {

	//  admin
	if IsAdmin(tenant) {
		return nil
	}

	//  seller company
	if tenant.CompanyID != nil &&
		*tenant.CompanyID == invoice.SellerCompanyID {
		return nil
	}

	//  buyer company
	if buyer.CompanyID != nil &&
		tenant.CompanyID != nil &&
		*buyer.CompanyID == *tenant.CompanyID {
		return nil
	}

	//  buyer individual
	if buyer.UserID != nil &&
		*buyer.UserID == tenant.UserID {
		return nil
	}

	return errorHandler.NewAppError(
		http.StatusForbidden,
		"No autorizado para ver esta invoice",
	)
}
