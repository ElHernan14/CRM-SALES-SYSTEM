package utils

import (
	"context"
	"fmt"
	"time"
)

func Trace(ctx context.Context, layer string) func(err error) {
	start := time.Now()

	return func(err error) {
		duration := time.Since(start)
		requestID := GetRequestID(ctx)

		status := "OK"
		if err != nil {
			status = "ERROR: " + err.Error()
		}

		// Fecha y hora completa
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		// Log de tracing más descriptivo
		fmt.Printf("[%s] Tracing → Capa=%s | id=%s | status=%s | duration=%s\n",
			timestamp, layer, requestID, status, duration)
	}
}
