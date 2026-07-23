package auth

import (
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, authController *AuthController) {
	// routes
	r.HandleFunc("/login", middleware.ErrorMiddleware(authController.Login)).Methods("POST")
	r.HandleFunc("/register/personal", middleware.ErrorMiddleware(authController.RegisterPersonal)).Methods("POST")
	r.HandleFunc("/register/business", middleware.ErrorMiddleware(authController.RegisterBusiness)).Methods("POST")
}
