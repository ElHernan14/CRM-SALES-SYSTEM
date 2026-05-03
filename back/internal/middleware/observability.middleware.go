package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	corecontext "crm-system-sales/internal/core/utils"

	"github.com/google/uuid"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// 🔥 middleware global
func ObservabilityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestID := uuid.New().String()

		// meter request_id en contexto
		ctx := context.WithValue(r.Context(), corecontext.RequestIDKey, requestID)

		// ejecutar request
		rw := &responseWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(rw, r.WithContext(ctx))

		// métricas + log
		duration := time.Since(start)
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		fmt.Printf("[%s] ObservabilityMiddleware → REQUEST id=%s method=%s path=%s status=%d duration=%fs\n",
			timestamp,
			requestID,
			r.Method,
			r.URL.Path,
			rw.statusCode,
			duration.Seconds(),
		)
		fmt.Printf("[%s] --------------------------------------------------------------------------------------------------------\n", timestamp)
	})
}
