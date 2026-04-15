package routes

import (
	"crm-system-sales/internal/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
}
