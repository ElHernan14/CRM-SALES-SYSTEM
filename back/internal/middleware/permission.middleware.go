package middleware

import (
	"net/http"

	"crm-system-sales/internal/utils"

	"crm-system-sales/internal/core"
)

func RequirePermission(permission string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			val := r.Context().Value(core.TenantContextKey)
			tenant, ok := val.(*core.TenantContext)
			if !ok || tenant == nil {
				http.Error(w, "Acceso denegado, usuario no autorizado por falta de permisosasd", http.StatusForbidden)
				return
			}

			if utils.HasPermission(tenant.Permissions, permission) {
				next(w, r)
				return
			}

			ctxTenant := core.GetTenant(r.Context())

			http.Error(w, "Acceso denegado, usuario no autorizado por falta de permisos."+ctxTenant.Email, http.StatusForbidden)
		}
	}
}
