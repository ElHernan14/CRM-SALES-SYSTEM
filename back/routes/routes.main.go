package routes

import (
	"database/sql"

	"crm-system-sales/internal/modules/auth"
	"crm-system-sales/internal/modules/client"
	"crm-system-sales/internal/modules/company"
	"crm-system-sales/internal/modules/inventory"
	invoiceController "crm-system-sales/internal/modules/invoice/controller"
	invoiceRepository "crm-system-sales/internal/modules/invoice/repository"
	invoiceRoutes "crm-system-sales/internal/modules/invoice/routes"
	invoiceService "crm-system-sales/internal/modules/invoice/service"
	invoiceitem "crm-system-sales/internal/modules/invoice_item"
	"crm-system-sales/internal/modules/product"
	"crm-system-sales/internal/modules/users"
	submitInvoiceWorkflow "crm-system-sales/internal/services/invoice_workflow/service"

	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

// AppContainer struct to hold controllers
type AppContainer struct {
	UserController        *users.UserController
	AuthController        *auth.AuthController
	ClientController      *client.ClientController
	CompanyController     *company.CompanyController
	ProductController     *product.ProductController
	InvoiceController     *invoiceController.InvoiceController
	InvoiceItemController *invoiceitem.InvoiceItemController
}

func SetupRoutes(r *mux.Router, db *sql.DB) {
	api := r.PathPrefix("/api").Subrouter()

	// auth (público)
	// Rutas públicas
	authRouter := api.PathPrefix("/auth").Subrouter()

	// rutas protegidos
	protected := api.NewRoute().Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// Repos
	userRepository := users.NewUserRepository(db)
	clientRepo := client.NewClientRepository(db)
	authRepo := auth.NewAuthRepository(db)
	companyRepo := company.NewCompanyRepository(db)
	productRepo := product.NewProductRepository(db)
	invoiceRepo := invoiceRepository.NewInvoiceRepository(db)
	invoiceItemRepo := invoiceitem.NewInvoiceItemRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	userService := users.NewUserService(userRepository)
	clientService := client.NewClientService(db, clientRepo, userRepository, authRepo)
	companyService := company.NewCompanyService(db, companyRepo, authRepo, userRepository)
	productService := product.NewProductService(db, productRepo)
	invoiceService := invoiceService.NewInvoiceService(db, invoiceRepo, clientRepo, companyRepo)
	inventoryService := inventory.NewInventoryService(productRepo)
	invoiceItemService := invoiceitem.NewInvoiceItemService(db, invoiceItemRepo, invoiceRepo, clientRepo, productRepo, inventoryService)
	// Workflow
	submitWorkflow := submitInvoiceWorkflow.NewSubmitInvoiceWorkflow(invoiceRepo, invoiceItemRepo, clientRepo, inventoryService, db)

	// Controllers
	authController := auth.NewAuthController(authService)
	userController := users.NewUserController(userService)
	clientController := client.NewClientController(clientService)
	companyController := company.NewCompanyController(companyService)
	productController := product.NewProductController(productService)
	invoiceController := invoiceController.NewInvoiceController(invoiceService, submitWorkflow)
	invoiceItemController := invoiceitem.NewInvoiceItemController(invoiceItemService)

	// Container para inyección de dependencias
	container := &AppContainer{
		UserController:        userController,
		AuthController:        authController,
		ClientController:      clientController,
		CompanyController:     companyController,
		ProductController:     productController,
		InvoiceController:     invoiceController,
		InvoiceItemController: invoiceItemController,
	}

	// Registrar rutas endpoints
	auth.RegisterAuthRoutes(authRouter, container.AuthController)
	users.RegisterUserRoutes(protected, container.UserController)
	client.RegisterClientRoutes(protected, container.ClientController)
	company.RegisterCompanyRoutes(protected, container.CompanyController)
	product.RegisterProductRoutes(protected, container.ProductController)
	invoiceRoutes.RegisterInvoiceRoutes(protected, container.InvoiceController)
	invoiceitem.RegisterInvoiceItemRoutes(protected, container.InvoiceItemController)
}
