package categories

import (
	"encoding/json"
	"net/http"
	"strconv"

	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/response"
)

type CategoriesController struct {
	service CategoriesService
}

func NewCategoriesController(service CategoriesService) *CategoriesController {
	return &CategoriesController{service: service}
}

func (c *CategoriesController) GetAll(w http.ResponseWriter, r *http.Request) error {
	res, err := c.service.GetAll(r.Context())
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *CategoriesController) GetProductTypes(w http.ResponseWriter, r *http.Request) error {
	var categoryID *int
	if v := r.URL.Query().Get("category_id"); v != "" {
		parsed, err := strconv.Atoi(v)
		if err != nil || parsed <= 0 {
			return errorHandler.NewAppError(http.StatusBadRequest, "category_id invalido")
		}
		categoryID = &parsed
	}

	res, err := c.service.GetProductTypes(r.Context(), categoryID)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(response.Success(res))
}
