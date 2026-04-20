package middleware

import (
	"net/http"

	"crm-system-sales/internal/utils"
)

func RequirePermission(permission string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			tenant, ok := r.Context().Value(TenantContextKey).(TenantContext)
			if !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			if utils.HasPermission(tenant.Permissions, permission) {
				next(w, r)
				return
			}

			http.Error(w, "Acceso denegado, usuario no autorizado por falta de permisos.", http.StatusForbidden)
		}
	}
}
