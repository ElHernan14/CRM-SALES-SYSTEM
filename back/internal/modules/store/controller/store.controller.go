package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	errorHandler "crm-system-sales/internal/core/error"
	helper "crm-system-sales/internal/core/helper"
	"crm-system-sales/internal/core/response"
	storedto "crm-system-sales/internal/modules/store/dto"
	service "crm-system-sales/internal/modules/store/service"

	"github.com/gorilla/mux"
)

type StoreController struct {
	Service service.StoreService
}

func NewStoreController(service service.StoreService) *StoreController {
	return &StoreController{Service: service}
}

func (c *StoreController) GetProducts(w http.ResponseWriter, r *http.Request) error {
	req, err := helper.ParseGetStoreProductsRequest(r)
	if err != nil {
		return err
	}
	res, err := c.Service.GetProducts(r.Context(), req)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *StoreController) GetProductByID(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "id invalido")
	}

	res, err := c.Service.GetProductByID(r.Context(), id)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *StoreController) GetCart(w http.ResponseWriter, r *http.Request) error {
	rawSellerID := r.URL.Query().Get("seller_company_id")
	if rawSellerID == "" {
		return errorHandler.NewAppError(http.StatusBadRequest, "seller_company_id es requerido")
	}
	sellerCompanyID, err := strconv.Atoi(rawSellerID)
	if err != nil || sellerCompanyID <= 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "seller_company_id invalido")
	}

	res, err := c.Service.GetCart(r.Context(), sellerCompanyID)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *StoreController) GetCarts(w http.ResponseWriter, r *http.Request) error {
	res, err := c.Service.GetCarts(r.Context())
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *StoreController) EnsureCart(w http.ResponseWriter, r *http.Request) error {
	var req storedto.EnsureCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errorHandler.NewAppError(http.StatusBadRequest, "json invalido")
	}
	res, err := c.Service.EnsureCart(r.Context(), &req)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *StoreController) CheckoutAll(w http.ResponseWriter, r *http.Request) error {
	req := &storedto.CheckoutAllRequest{}
	if r.Body != nil && r.ContentLength > 0 {
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "json invalido")
		}
	}
	res, err := c.Service.CheckoutAll(r.Context(), req)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *StoreController) Checkout(w http.ResponseWriter, r *http.Request) error {
	req := &storedto.CheckoutRequest{}
	if r.Body != nil && r.ContentLength > 0 {
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "json invalido")
		}
	}
	res, err := c.Service.Checkout(r.Context(), req)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(response.Success(res))
}
