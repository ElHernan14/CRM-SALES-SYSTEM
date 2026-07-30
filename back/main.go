package main

import (
	"log"
	"net/http"

	"crm-system-sales/internal/config"
	"crm-system-sales/internal/controllers"
	"crm-system-sales/internal/db"
	"crm-system-sales/internal/middleware"
	"crm-system-sales/routes"

	_ "crm-system-sales/docs"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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

	healthHandler := controllers.NewHealthHandler(database)

	r := mux.NewRouter()
	r.Use(middleware.RecoverMiddleware)
	r.Use(middleware.RequestIDMiddleware)
	r.Use(middleware.ObservabilityMiddleware)

	r.HandleFunc("/health", healthHandler.Check).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.PathPrefix("/uploads/").Handler(
		http.StripPrefix("/uploads/", http.FileServer(http.Dir(cfg.UploadDir))),
	)

	api := r.PathPrefix("/api").Subrouter()
	routes.SetupRoutes(api, database, cfg)

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Server running on :%s", cfg.ServerPort)

	cors := handlers.CORS(
		handlers.AllowedOrigins(cfg.CORSOrigins),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, cors(r)))
}
