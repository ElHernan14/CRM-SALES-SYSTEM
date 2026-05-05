package access

import (
	"crm-system-sales/internal/constants"
	authcore "crm-system-sales/internal/core/auth"
	errorHandler "crm-system-sales/internal/core/error"
	tenant "crm-system-sales/internal/core/tenant"
	"net/http"
)

func ResolveClientScope(tenant *tenant.TenantContext, reqCompanyID *int) (*int, error) {
	var err error
	if authcore.HasPermission(tenant.Permissions, constants.ClientViewAll) {
		return reqCompanyID, nil
	}
	if authcore.HasPermission(tenant.Permissions, constants.ClientViewCompany) {
		if tenant.CompanyID == nil {
			err = errorHandler.NewAppError(http.StatusForbidden, "company_id not found")
			return nil, err
		}
		// evitar bypass por query
		if reqCompanyID != nil && *reqCompanyID != *tenant.CompanyID {
			err = errorHandler.NewAppError(http.StatusForbidden, "company_id does not match")
			return nil, err
		}

		return tenant.CompanyID, nil
	}
	err = errorHandler.NewAppError(http.StatusForbidden, "access denied")
	return nil, err
}
