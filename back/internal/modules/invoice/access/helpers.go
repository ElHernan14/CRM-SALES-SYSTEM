package invoiceaccess

import (
	constants "crm-system-sales/internal/constants"
	tenantctx "crm-system-sales/internal/core/tenant"
)

func HasRole(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

func IsAdmin(t *tenantctx.TenantContext) bool {
	return HasRole(t.Roles, constants.RoleAdmin)
}
