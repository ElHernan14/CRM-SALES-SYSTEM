package access

import (
	constants "crm-system-sales/internal/constants"
	errorHandler "crm-system-sales/internal/core/error"
	tenant "crm-system-sales/internal/core/tenant"
	"net/http"
)

func ResolveCompanyScope(tenant *tenant.TenantContext, companyID *int) (*int, error) {
	var err error
	if hasRol(tenant.Roles, constants.RoleAdmin) {
		return companyID, nil
	}
	if tenant.CompanyID == nil || *tenant.CompanyID != *companyID {
		err = errorHandler.NewAppError(http.StatusForbidden, "Acceso denegado a empresa.")
		return nil, err
	}
	return companyID, nil
}

func hasRol(roles []string, role string) bool {
	for _, p := range roles {
		if p == role {
			return true
		}
	}
	return false
}
