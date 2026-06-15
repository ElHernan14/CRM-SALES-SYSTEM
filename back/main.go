// @title CRM System Sales API
// @version 1.0
// @description Multi-tenant ERP & Commerce SaaS API
// @BasePath /api
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http
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

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	/*"golang.org/x/crypto/bcrypt"*/)

//
//	-- Departamentos de Ventas, Servicio de Atención al Cliente, Marketing y Medios de Comunicación (historial de compras del cliente,
// 	el estado de sus pedidos, los problemas pendientes del servicio de atención al cliente, etc),
// 	y un panel de administración para gestionar productos, pedidos, clientes, etc. cada uno de estos con su propio conjunto de rutas y controladores.
//
//	- Gestión de Clientes: CRUD de clientes, historial de compras, preferencias, etc.
//	- Gestión de Productos: CRUD de productos, categorías, precios, etc.
//	- Gestión de Pedidos: CRUD de pedidos, seguimiento, estado, etc.
//	- Gestión de Inventario: CRUD de inventario, niveles de stock, alertas, etc.
//	- Gestión de Usuarios y Roles: CRUD de usuarios, asignación de roles, permisos, etc.

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

	// Swagger route
	r.PathPrefix("/swagger/").Handler(
		httpSwagger.WrapHandler,
	)

	//Tracing middleware
	r.Use(middleware.ObservabilityMiddleware)

	routes.SetupRoutes(r, database)

	// hash, _ := bcrypt.GenerateFromPassword([]byte("asd123"), bcrypt.DefaultCost)
	// log.Println(string(hash))

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
