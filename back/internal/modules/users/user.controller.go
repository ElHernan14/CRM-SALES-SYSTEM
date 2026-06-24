package users

import (
	"encoding/json"
	"net/http"

	tenantHelper "crm-system-sales/internal/core/tenant"
)

type UserController struct {
	UserService UserService
}

func NewUserController(UserService UserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

func (c *UserController) Me(w http.ResponseWriter, r *http.Request) {
	tenant := tenantHelper.GetTenant(r.Context())

	response := map[string]interface{}{
		"userID":      tenant.UserID,
		"email":       tenant.Email,
		"company_id":  tenant.CompanyID,
		"client_id":   tenant.ClientID,
		"roles":       tenant.Roles,
		"permissions": tenant.Permissions,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

	users := []map[string]string{
		{"id": "1", "name": "Hernan"},
		{"id": "2", "name": "Juan"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	users := []map[string]string{
		{"id": "1", "name": "Hernan"},
		{"id": "2", "name": "Juan"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
