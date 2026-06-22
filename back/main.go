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

	// =========================
	// ROOT ROUTER
	// =========================
	r := mux.NewRouter()

	// =========================
	// GLOBAL MIDDLEWARES
	// =========================
	r.Use(middleware.RecoverMiddleware)
	r.Use(middleware.RequestIDMiddleware)
	r.Use(middleware.ObservabilityMiddleware)

	// =========================
	// OPTIONS GLOBAL HANDLER
	// (IMPORTANTE para preflight)
	// =========================
	// r.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	// 		w.Header().Set("Vary", "Origin")
	// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// 		if r.Method == "OPTIONS" {
	// 			w.WriteHeader(204)
	// 			return
	// 		}

	// 		next.ServeHTTP(w, r)
	// 	})
	// })

	// =========================
	// HEALTH + SWAGGER
	// =========================
	r.HandleFunc("/health", healthHandler.Check).Methods("GET")

	r.PathPrefix("/swagger/").Handler(
		httpSwagger.WrapHandler,
	)

	// =========================
	// API SUBROUTER (OK ACA)
	// =========================
	api := r.PathPrefix("/api").Subrouter()

	// rutas de la app
	routes.SetupRoutes(api, database)
	// r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	// 	path, _ := route.GetPathTemplate()
	// 	methods, _ := route.GetMethods()
	// 	log.Println("ROUTE:", path, methods)
	// 	return nil
	// })

	// =========================
	// START SERVER
	// =========================
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Server running on :8081")

	// r.HandleFunc("/api/auth/login", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("MANUAL OPTIONS HIT")

	// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	// 	w.WriteHeader(http.StatusNoContent)
	// }).Methods("OPTIONS")

	cors := handlers.CORS(
		handlers.AllowedOrigins(
			[]string{"http://localhost:5173"},
		),
		handlers.AllowedMethods(
			[]string{
				"GET",
				"POST",
				"PUT",
				"PATCH",
				"DELETE",
				"OPTIONS",
			},
		),
		handlers.AllowedHeaders(
			[]string{
				"Content-Type",
				"Authorization",
			},
		),
	)

	log.Fatal(
		http.ListenAndServe(
			":8081",
			cors(r),
		),
	)
}
