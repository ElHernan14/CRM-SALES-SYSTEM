package invoiceaccess

import (
	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"

	clientModel "crm-system-sales/internal/models/client"

	"net/http"
)

func CanCreateInvoice(
	tenant *tenantctx.TenantContext,
	buyer *clientModel.Client,
	sellerCompanyID int,
) error {

	if IsAdmin(tenant) {
		return nil
	}

	// individual user
	if tenant.CompanyID == nil {

		if buyer.UserID == nil ||
			*buyer.UserID != tenant.UserID {

			return errorHandler.NewAppError(
				http.StatusForbidden,
				"No autorizado para crear invoice",
			)
		}

		return nil
	}

	// empresa creando compra
	if buyer.CompanyID != nil &&
		*buyer.CompanyID == *tenant.CompanyID {
		return nil
	}

	// empresa creando venta
	if *tenant.CompanyID == sellerCompanyID {
		return nil
	}

	return errorHandler.NewAppError(
		http.StatusForbidden,
		"No autorizado para crear invoice",
	)
}
