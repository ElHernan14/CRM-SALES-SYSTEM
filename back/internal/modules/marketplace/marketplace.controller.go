package marketplace

import (
	"crm-system-sales/internal/core/helper"
	"crm-system-sales/internal/core/response"
	"encoding/json"
	"net/http"
	"strconv"

	errorHandler "crm-system-sales/internal/core/error"

	"github.com/gorilla/mux"
)

type MarketplaceController struct {
	service MarketplaceService
}

func NewMarketplaceController(service MarketplaceService) *MarketplaceController {
	return &MarketplaceController{service: service}
}

func (c *MarketplaceController) GetSuppliers(w http.ResponseWriter, r *http.Request) error {
	req, err := helper.ParseGetSuppliersRequest(r)
	if err != nil {
		return err
	}

	res, err := c.service.GetSuppliers(r.Context(), req)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *MarketplaceController) GetSupplierByID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "supplier_id invalido")
	}

	res, err := c.service.GetSupplierByID(r.Context(), id)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}
