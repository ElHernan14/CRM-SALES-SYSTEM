package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	authcore "crm-system-sales/internal/core/auth"
	"crm-system-sales/internal/core/response"

	tenantHelper "crm-system-sales/internal/core/tenant"
)

func RequirePermission(permission string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			tenant := tenantHelper.GetTenant(r.Context())
			if tenant == nil {
				log.Printf("Tenant no encontrado en el contexto")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(response.Error(http.StatusForbidden, "Acceso denegado, usuario no autorizado por falta de permiso"))
				return
			}

			if authcore.HasPermission(tenant.Permissions, permission) {
				next(w, r)
				return
			}

			ctxTenant := tenantHelper.GetTenant(r.Context())

			log.Printf("Acceso denegado, usuario no autorizado por falta de permisos: %s", ctxTenant.Email)
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(response.Error(http.StatusForbidden, "Acceso denegado, usuario no autorizado por falta de permisos. "+ctxTenant.Email))
		}
	}
}
