package auth

import (
	"database/sql"

	"crm-system-sales/internal/modules/users"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, db *sql.DB) {
	// dependency injection
	userRepository := users.NewUserRepository(db)
	authService := NewAuthService(userRepository)
	authController := NewAuthController(authService)

	// routes
	r.HandleFunc("/login", authController.Login).Methods("POST")
}
