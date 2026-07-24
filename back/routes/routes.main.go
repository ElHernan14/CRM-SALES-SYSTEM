package routes

import (
	"database/sql"

	"crm-system-sales/internal/config"
	"crm-system-sales/internal/core/files"
	"crm-system-sales/internal/middleware"
	"crm-system-sales/internal/modules/auth"
	"crm-system-sales/internal/modules/categories"
	categorycompany "crm-system-sales/internal/modules/category_company"
	categoryproduct "crm-system-sales/internal/modules/category_product"
	"crm-system-sales/internal/modules/client"
	"crm-system-sales/internal/modules/company"
	"crm-system-sales/internal/modules/dashboard"
	"crm-system-sales/internal/modules/inventory"
	invoiceController "crm-system-sales/internal/modules/invoice/controller"
	invoiceRepository "crm-system-sales/internal/modules/invoice/repository"
	invoiceRoutes "crm-system-sales/internal/modules/invoice/routes"
	invoiceService "crm-system-sales/internal/modules/invoice/service"
	invoiceitem "crm-system-sales/internal/modules/invoice_item"
	invoicepayment "crm-system-sales/internal/modules/invoice_payment"
	invoicepaymentrepo "crm-system-sales/internal/modules/invoice_payment/repository"
	"crm-system-sales/internal/modules/marketplace"
	"crm-system-sales/internal/modules/product"
	producttype "crm-system-sales/internal/modules/product_type"
	storecontroller "crm-system-sales/internal/modules/store/controller"
	storeroutes "crm-system-sales/internal/modules/store/routes"
	storeservice "crm-system-sales/internal/modules/store/service"
	"crm-system-sales/internal/modules/users"
	dashboardoverview "crm-system-sales/internal/services/dashboard_overview/service"
	paymentworkflow "crm-system-sales/internal/services/invoice_workflow/service"
	submitInvoiceWorkflow "crm-system-sales/internal/services/invoice_workflow/service"

	"github.com/gorilla/mux"
)

// AppContainer struct to hold controllers
type AppContainer struct {
	UserController           *users.UserController
	AuthController           *auth.AuthController
	ClientController         *client.ClientController
	CompanyController        *company.CompanyController
	ProductController        *product.ProductController
	InvoiceController        *invoiceController.InvoiceController
	InvoiceItemController    *invoiceitem.InvoiceItemController
	InvoicePaymentController *invoicepayment.InvoicePaymentController
	StoreController          *storecontroller.StoreController
	MarketplaceController    *marketplace.MarketplaceController
	CategoriesController     *categories.CategoriesController
	DashboardController      *dashboard.DashboardController
}

func SetupRoutes(r *mux.Router, db *sql.DB, cfg config.Config) {
	authRouter := r.PathPrefix("/auth").Subrouter()

	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// Repos
	userRepository := users.NewUserRepository(db)
	clientRepo := client.NewClientRepository(db)
	authRepo := auth.NewAuthRepository(db)
	companyRepo := company.NewCompanyRepository(db)
	categoryCompanyRepo := categorycompany.NewCategoryCompanyRepository(db)
	categoryProductRepo := categoryproduct.NewCategoryProductRepository(db)
	productTypeRepo := producttype.NewProductTypeRepository(db)
	productRepo := product.NewProductRepository(db)
	invoiceRepo := invoiceRepository.NewInvoiceRepository(db)
	invoiceItemRepo := invoiceitem.NewInvoiceItemRepository(db)
	invoicePaymentRepo := invoicepaymentrepo.NewInvoicePaymentRepository(db)
	marketplaceRepo := marketplace.NewMarketplaceRepository(db)
	dashboardRepo := dashboard.NewDashboardRepository(db)

	// Shared services
	imageStorage := files.NewLocalImageStorage(cfg.UploadDir)

	// Services
	authService := auth.NewAuthService(db, userRepository, clientRepo, companyRepo, authRepo, categoryCompanyRepo)
	userService := users.NewUserService(userRepository)
	clientService := client.NewClientService(db, clientRepo, userRepository, authRepo)
	companyService := company.NewCompanyService(db, companyRepo, clientRepo, authRepo, userRepository, categoryCompanyRepo, imageStorage)
	productService := product.NewProductService(db, productRepo, categoryProductRepo, productTypeRepo, imageStorage)
	inventoryService := inventory.NewInventoryService(productRepo)
	invoiceItemService := invoiceitem.NewInvoiceItemService(db, invoiceItemRepo, invoiceRepo, clientRepo, productRepo, inventoryService)
	submitWorkflow := submitInvoiceWorkflow.NewSubmitInvoiceWorkflow(invoiceRepo, invoiceItemRepo, clientRepo, inventoryService, db)
	paymentWorkflow := paymentworkflow.NewPayInvoiceWorkflow(invoiceItemRepo, inventoryService)
	invoiceService := invoiceService.NewInvoiceService(db, invoiceRepo, clientRepo, companyRepo, invoiceItemRepo, inventoryService, invoicePaymentRepo, paymentWorkflow)
	invoicePaymentService := invoicepayment.NewInvoicePaymentService(db, invoiceRepo, clientRepo, invoicePaymentRepo)
	storeService := storeservice.NewStoreService(db, productRepo, invoiceRepo, invoiceItemRepo, companyRepo, submitWorkflow)
	marketplaceService := marketplace.NewMarketplaceService(marketplaceRepo)
	categoriesService := categories.NewCategoriesService(categoryProductRepo, categoryCompanyRepo, productTypeRepo)
	dashboardService := dashboardoverview.NewDashboardOverviewService(dashboardRepo)

	// Controllers
	authController := auth.NewAuthController(authService)
	userController := users.NewUserController(userService)
	clientController := client.NewClientController(clientService)
	companyController := company.NewCompanyController(companyService)
	productController := product.NewProductController(productService)
	invoiceController := invoiceController.NewInvoiceController(invoiceService, submitWorkflow)
	invoiceItemController := invoiceitem.NewInvoiceItemController(invoiceItemService)
	invoicePaymentController := invoicepayment.NewInvoicePaymentController(invoicePaymentService)
	storeController := storecontroller.NewStoreController(storeService)
	marketplaceController := marketplace.NewMarketplaceController(marketplaceService)
	categoriesController := categories.NewCategoriesController(categoriesService)
	dashboardController := dashboard.NewDashboardController(dashboardService)

	container := &AppContainer{
		UserController:           userController,
		AuthController:           authController,
		ClientController:         clientController,
		CompanyController:        companyController,
		ProductController:        productController,
		InvoiceController:        invoiceController,
		InvoiceItemController:    invoiceItemController,
		InvoicePaymentController: invoicePaymentController,
		StoreController:          storeController,
		MarketplaceController:    marketplaceController,
		CategoriesController:     categoriesController,
		DashboardController:      dashboardController,
	}

	auth.RegisterAuthRoutes(authRouter, container.AuthController)
	storeroutes.RegisterPublicStoreRoutes(r, container.StoreController)
	categories.RegisterCategoriesRoutes(r, container.CategoriesController)

	users.RegisterUserRoutes(protected, container.UserController)
	client.RegisterClientRoutes(protected, container.ClientController)
	company.RegisterCompanyRoutes(protected, container.CompanyController)
	product.RegisterProductRoutes(protected, container.ProductController)
	invoiceRoutes.RegisterInvoiceRoutes(protected, container.InvoiceController)
	invoiceitem.RegisterInvoiceItemRoutes(protected, container.InvoiceItemController)
	invoicepayment.RegisterInvoicePaymentRoutes(protected, container.InvoicePaymentController)
	storeroutes.RegisterInvoiceItemRoutes(protected, container.StoreController)
	marketplace.RegisterMarketplaceRoutes(protected, container.MarketplaceController)
	dashboard.RegisterDashboardRoutes(protected, container.DashboardController)
}
