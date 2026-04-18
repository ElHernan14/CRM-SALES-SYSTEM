package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func RequirePermission(permission string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(UserContextKey).(jwt.MapClaims)
			if !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			perms, ok := claims["permissions"]
			if !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			permSlice, ok := perms.([]interface{})
			if !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			for _, p := range permSlice {
				if p.(string) == permission {
					next(w, r)
					return
				}
			}

			http.Error(w, "Acceso denegado, usuario no autorizado por falta de permisos.", http.StatusForbidden)
		}
	}
}
