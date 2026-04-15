package routes

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	// subrutas
	RegisterUserRoutes(api)
}
