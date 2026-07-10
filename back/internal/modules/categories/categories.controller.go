package categories

import (
	"encoding/json"
	"net/http"

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
