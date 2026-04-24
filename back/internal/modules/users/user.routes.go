package users

import (
	"database/sql"

	"github.com/gorilla/mux"

	"crm-system-sales/internal/middleware"
)

func RegisterUserRoutes(r *mux.Router, db *sql.DB) {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userController := NewUserController(userService)

	r.HandleFunc("/me", userController.Me).Methods("GET")
	r.HandleFunc("/users", userController.GetUsers).Methods("GET")
	r.HandleFunc("/create", middleware.RequirePermission("user:create")(userController.CreateUser)).Methods("POST")
}
