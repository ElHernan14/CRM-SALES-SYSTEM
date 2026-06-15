package controller

import (
	helper "crm-system-sales/internal/core/helper"
	"crm-system-sales/internal/core/response"
	service "crm-system-sales/internal/modules/store/service"
	"encoding/json"
	"net/http"
)

type StoreController struct {
	Service service.StoreService
}

func NewStoreController(service service.StoreService) *StoreController {
	return &StoreController{
		Service: service,
	}
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
		req,
	)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(
		response.Success(res),
	)
}

func (c *StoreController) Checkout(
	w http.ResponseWriter,
	r *http.Request,
) error {

	res, err := c.Service.Checkout(
		r.Context(),
	)

	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(
		response.Success(res),
	)
}
