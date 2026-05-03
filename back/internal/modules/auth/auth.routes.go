package auth

import (
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, authController *AuthController) {
	// routes
	r.HandleFunc("/login", middleware.ErrorMiddleware(authController.Login)).Methods("POST")
}
