package controller

import (
	helper "crm-system-sales/internal/core/helper"
	"crm-system-sales/internal/core/response"
	"encoding/json"
	"net/http"
)

type StoreController struct {
	Service interface{}
}

func NewStoreController() *StoreController {
	return &StoreController{}
}

func (c *StoreController) GetProducts(
	w http.ResponseWriter,
	r *http.Request,
) error {

	req, err := helper.ParseGetStoreProductsRequest(r)
	if err != nil {
		return err
	}

	res, err := c.Service.GetProducts(
		r.Context(),
		&req,
	)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(
		response.Success(res),
	)
}
