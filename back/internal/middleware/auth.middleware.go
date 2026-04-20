package middleware

import (
	"context"
	"net/http"
	"strings"

	"crm-system-sales/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

type TenantContext struct {
	UserID      int
	Email       string
	CompanyID   *int
	Permissions []string
}

type contextKey string

const TenantContextKey contextKey = "tenant"

var jwtSecret = []byte("super_secret_key")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &utils.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		tenant := TenantContext{
			UserID:      claims.UserID,
			CompanyID:   claims.CompanyID,
			Permissions: claims.Permissions,
			Email:       claims.Email,
		}

		ctx := context.WithValue(r.Context(), TenantContextKey, tenant)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
