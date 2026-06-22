package middleware

import (
	"log"
	"net/http"
	"time"

	corecontext "crm-system-sales/internal/core/utils"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// middleware global
func ObservabilityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		rw := &responseWriter{ResponseWriter: w, statusCode: 200}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		reqID := r.Context().Value(corecontext.RequestIDKey)

		log.Printf(
			"[OBS] id=%v method=%s path=%s status=%d duration=%s",
			reqID,
			r.Method,
			r.URL.Path,
			rw.statusCode,
			duration,
		)
	})
}
