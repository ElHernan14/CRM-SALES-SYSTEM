package routes

import (
	"database/sql"

	"crm-system-sales/internal/modules/auth"
	"crm-system-sales/internal/modules/users"

	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, db *sql.DB) {
	api := r.PathPrefix("/api").Subrouter()

	//auth (público)
	authRouter := api.PathPrefix("/auth").Subrouter()

	auth.RegisterAuthRoutes(authRouter, db)

	// rutas protegidos
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	users.RegisterUserRoutes(protected, db)
}
