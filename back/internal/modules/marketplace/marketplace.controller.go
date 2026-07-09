package marketplace

import (
	"crm-system-sales/internal/core/helper"
	"crm-system-sales/internal/core/response"
	"encoding/json"
	"net/http"
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
