package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	authcore "crm-system-sales/internal/core/auth"
	tenantHelper "crm-system-sales/internal/core/tenant"
)

var jwtSecret = []byte("super_secret_key")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &authcore.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		tenant := &tenantHelper.TenantContext{
			UserID:      claims.UserID,
			CompanyID:   claims.CompanyID,
			Email:       claims.Email,
			Roles:       *claims.Roles,
			Permissions: *claims.Permissions,
		}

		ctx := context.WithValue(r.Context(), tenantHelper.TenantContextKey, tenant)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
