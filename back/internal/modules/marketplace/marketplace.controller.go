package marketplace

import (
	"crm-system-sales/internal/core/helper"
	"crm-system-sales/internal/core/response"
	validatorx "crm-system-sales/internal/core/validator"
	marketplacedto "crm-system-sales/internal/modules/marketplace/dto"
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

func (c *MarketplaceController) EnsureCart(w http.ResponseWriter, r *http.Request) error {
	var req marketplacedto.EnsureMarketplaceCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errorHandler.NewAppError(http.StatusBadRequest, "json invalido")
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	res, err := c.service.EnsureCart(r.Context(), &req)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *MarketplaceController) GetCart(w http.ResponseWriter, r *http.Request) error {
	rawSellerID := r.URL.Query().Get("seller_company_id")
	if rawSellerID == "" {
		return errorHandler.NewAppError(http.StatusBadRequest, "seller_company_id es requerido")
	}

	sellerCompanyID, err := strconv.Atoi(rawSellerID)
	if err != nil || sellerCompanyID <= 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "seller_company_id invalido")
	}

	res, err := c.service.GetCart(r.Context(), sellerCompanyID)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *MarketplaceController) Checkout(w http.ResponseWriter, r *http.Request) error {
	req := &marketplacedto.MarketplaceCheckoutRequest{}
	if r.Body != nil && r.ContentLength > 0 {
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "json invalido")
		}
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	res, err := c.service.Checkout(r.Context(), req)
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
