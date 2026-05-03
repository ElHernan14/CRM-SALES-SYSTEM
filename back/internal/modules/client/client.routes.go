package client

import (
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterClientRoutes(r *mux.Router, controller *ClientController) {
	clients := r.PathPrefix("/clients").Subrouter()
	clients.HandleFunc("/create", middleware.RequirePermission("client:create")(middleware.ErrorMiddleware(controller.Create))).Methods("POST")
	clients.HandleFunc("/clients", middleware.RequirePermission("client:view")(middleware.ErrorMiddleware(controller.GetClients))).Methods("GET")
}
