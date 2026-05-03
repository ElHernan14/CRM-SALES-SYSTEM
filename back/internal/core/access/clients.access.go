package access

import (
	"crm-system-sales/internal/constants"
	"crm-system-sales/internal/core"
	errorHandler "crm-system-sales/internal/core/error"
	utilsx "crm-system-sales/internal/utils"
)

func ResolveClientScope(tenant *core.TenantContext, reqCompanyID *int) (*int, error) {
	if utilsx.HasPermission(tenant.Permissions, constants.ClientViewAll) {
		return reqCompanyID, nil
	}
	if utilsx.HasPermission(tenant.Permissions, constants.ClientViewCompany) {
		if tenant.CompanyID == nil {
			return nil, errorHandler.ErrForbidden
		}
		// evitar bypass por query
		if reqCompanyID != nil && *reqCompanyID != *tenant.CompanyID {
			return nil, errorHandler.ErrForbidden
		}

		return tenant.CompanyID, nil
	}
	return nil, errorHandler.ErrForbidden
}
