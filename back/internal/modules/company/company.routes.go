package company

import (
	constants "crm-system-sales/internal/constants"
	"crm-system-sales/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterCompanyRoutes(r *mux.Router, controller *CompanyController) {
	companies := r.PathPrefix("/company").Subrouter()
	companies.HandleFunc("/me", middleware.ErrorMiddleware(controller.GetMyCompany)).Methods("GET")
	companies.HandleFunc("/me/logo", middleware.RequirePermission(constants.CompanyUpdate)(middleware.ErrorMiddleware(controller.UploadLogo))).Methods("POST")
	companies.HandleFunc("/me/cover-image", middleware.RequirePermission(constants.CompanyUpdate)(middleware.ErrorMiddleware(controller.UploadCoverImage))).Methods("POST")
	companies.HandleFunc("/create", middleware.RequirePermission(constants.CompanyCreate)(middleware.ErrorMiddleware(controller.Create))).Methods("POST")
	companies.HandleFunc("/companies", middleware.RequirePermission(constants.CompanyView)(middleware.ErrorMiddleware(controller.GetCompanies))).Methods("GET")
	companies.HandleFunc("/companies/{id}", middleware.RequirePermission(constants.CompanyView)(middleware.ErrorMiddleware(controller.GetCompanyByID))).Methods("GET")
}
