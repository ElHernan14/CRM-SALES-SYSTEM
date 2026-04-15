package main

import (
	"log"
	"net/http"

	"crm-system-sales/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.SetupRoutes(r)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
