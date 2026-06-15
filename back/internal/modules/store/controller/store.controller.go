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

// GetProducts godoc
//
// @Summary Store catalog
// @Description Returns products available for ecommerce purchase
// @Tags Store
// @Produce json
// @Security BearerAuth
//
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param name query string false "Product name"
// @Param type query string false "Product type"
// @Param sort_column query string false "Sort column"
// @Param order query string false "Sort order"
//
// @Success 200 {object} docs.GetStoreProductsSuccessResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
//
// @Router /store/products [get]
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

// Checkout godoc
//
// @Summary Checkout cart
// @Description Submits the user's draft invoice and reserves stock
// @Tags Store
// @Accept json
// @Produce json
// @Security BearerAuth
//
// @Success 200 {object} docs.CheckoutSuccessResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
//
// @Router /store/checkout [post]
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
