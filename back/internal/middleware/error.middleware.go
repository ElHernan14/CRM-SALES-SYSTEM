package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/response"
)

// handler que devuelve error
type AppHandler func(w http.ResponseWriter, r *http.Request) error

func ErrorMiddleware(next AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("METHOD LLEGADO", r.Method)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		err := next(w, r)
		if err == nil {
			return
		}

		// timestamp completo
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		// error controlado
		if appErr, ok := err.(errorHandler.AppError); ok {
			fmt.Printf("[%s] ErrorMiddleware → Status=%d, Message=%s\n",
				timestamp, appErr.Code, appErr.Message)
			w.WriteHeader(appErr.Code)
			json.NewEncoder(w).Encode(response.Error(appErr.Code, appErr.Message))
			return
		}

		// error inesperado
		fmt.Printf("[%s] ErrorMiddleware → Error inesperado: %v\n",
			timestamp, err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.Error(http.StatusInternalServerError, errorHandler.ErrInternal.Error()))
	}
}
