package client

import (
	constants "crm-system-sales/internal/constants"
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterClientRoutes(r *mux.Router, controller *ClientController) {
	clients := r.PathPrefix("/clients").Subrouter()
	clients.HandleFunc("/create", middleware.RequirePermission(constants.ClientCreate)(middleware.ErrorMiddleware(controller.Create))).Methods("POST")
	clients.HandleFunc("/clients", middleware.RequirePermission(constants.ClientView)(middleware.ErrorMiddleware(controller.GetClients))).Methods("GET")
	clients.HandleFunc("/clients/{id}", middleware.RequirePermission(constants.ClientView)(middleware.ErrorMiddleware(controller.GetClientByID))).Methods("GET")
	clients.HandleFunc("/clients/{id}", middleware.RequirePermission(constants.ClientUpdate)(middleware.ErrorMiddleware(controller.UpdateClient))).Methods("PATCH")
	clients.HandleFunc("/clients/{id}", middleware.RequirePermission(constants.ClientDelete)(middleware.ErrorMiddleware(controller.DeleteClient))).Methods("DELETE")
}
