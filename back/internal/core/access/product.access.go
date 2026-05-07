package access

import (
	constants "crm-system-sales/internal/constants"
	errorHandler "crm-system-sales/internal/core/error"
	tenantctx "crm-system-sales/internal/core/tenant"
	"net/http"
)

func ResolveCreateProductCompanyID(
	tenant *tenantctx.TenantContext,
	reqCompanyID *int,
) (*int, error) {
	if hasRole(tenant.Roles, constants.RoleAdmin) {
		if reqCompanyID == nil {
			return nil, errorHandler.NewAppError(
				http.StatusBadRequest,
				"company_id requerido para admin",
			)
		}
		return reqCompanyID, nil
	}

	if tenant.CompanyID != nil {
		return tenant.CompanyID, nil
	}

	return nil, errorHandler.NewAppError(http.StatusForbidden, "Usuario sin empresa asignada")
}

func ResolveGetProductsCompanyID(
	tenant *tenantctx.TenantContext,
	reqCompanyID *int,
) (*int, error) {
	if hasRole(tenant.Roles, constants.RoleAdmin) {
		return reqCompanyID, nil // nil = todos
	}

	if tenant.CompanyID != nil {
		if reqCompanyID != nil && *reqCompanyID != *tenant.CompanyID {
			return nil, errorHandler.NewAppError(
				http.StatusForbidden,
				"No se permite consultar productos de otra empresa",
			)
		}
		return tenant.CompanyID, nil
	}

	return nil, nil
}

func hasRole(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
