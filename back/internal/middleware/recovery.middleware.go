package middleware

import (
	"crm-system-sales/internal/core/response"
	"encoding/json"
	"log"
	"net/http"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v", err)

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response.Error(
					500,
					"internal server error",
				))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
