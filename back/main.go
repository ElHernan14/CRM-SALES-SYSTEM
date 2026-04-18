package main

import (
	"log"
	"net/http"

	"crm-system-sales/internal/config"
	"crm-system-sales/internal/controllers"
	"crm-system-sales/internal/db"

	"crm-system-sales/routes"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	database := db.NewPostgresDB(db.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
		SSLMode:  cfg.DBSSLMode,
	})

	_ = database

	healthHandler := controllers.NewHealthHandler(database)

	// router
	r := mux.NewRouter()

	// Health check route
	r.HandleFunc("/health", healthHandler.Check).Methods("GET")

	routes.SetupRoutes(r, database)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
