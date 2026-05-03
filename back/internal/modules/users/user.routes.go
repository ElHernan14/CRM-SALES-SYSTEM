package users

import (
	"github.com/gorilla/mux"

	"crm-system-sales/internal/middleware"
)

func RegisterUserRoutes(r *mux.Router, userController *UserController) {
	r.HandleFunc("/me", userController.Me).Methods("GET")
	r.HandleFunc("/users", userController.GetUsers).Methods("GET")
	r.HandleFunc("/create", middleware.RequirePermission("user:create")(userController.CreateUser)).Methods("POST")
}
