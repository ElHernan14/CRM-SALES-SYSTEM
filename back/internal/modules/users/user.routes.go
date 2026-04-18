package users

import (
	"database/sql"

	"github.com/gorilla/mux"

	"crm-system-sales/internal/middleware"
)

func RegisterUserRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/create", middleware.RequirePermission("user:create")(CreateUser)).Methods("POST")
}
