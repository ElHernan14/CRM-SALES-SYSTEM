package invoiceitem

import (
	errorHandler "crm-system-sales/internal/core/error"
	helper "crm-system-sales/internal/core/helper"
	"crm-system-sales/internal/core/response"
	validatorx "crm-system-sales/internal/core/validator"
	invoiceitemdto "crm-system-sales/internal/modules/invoice_item/dto"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InvoiceItemController struct {
	service InvoiceItemService
}

func NewInvoiceItemController(service InvoiceItemService) *InvoiceItemController {
	return &InvoiceItemController{service: service}
}

// Create godoc
//
// @Summary Add item to invoice
// @Description Adds a product to a draft invoice or increases quantity if already exists
// @Tags Invoice Items
// @Accept json
// @Produce json
// @Security BearerAuth
//
// @Param id path int true "Invoice ID"
// @Param request body invoiceitemdto.CreateInvoiceItemRequest true "Invoice Item"
//
// @Success 200 {object} docs.CreateInvoiceItemSuccessResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
//
// @Router /invoice/{id}/items [post]
func (c *InvoiceItemController) Create(
	w http.ResponseWriter,
	r *http.Request,
) error {

	params := mux.Vars(r)

	invoiceID, err := strconv.Atoi(params["id"])
	if err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"invoice id inválido",
		)
	}

	var req invoiceitemdto.CreateInvoiceItemRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"json inválido",
		)
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			msg,
		)
	}

	res, err := c.service.Create(
		r.Context(),
		invoiceID,
		&req,
	)

	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(
		response.Success(res),
	)

	return nil
}

// Update godoc
//
// @Summary Update invoice item
// @Description Updates item quantity and recalculates subtotal
// @Tags Invoice Items
// @Accept json
// @Produce json
// @Security BearerAuth
//
// @Param id path int true "Invoice ID"
// @Param itemId path int true "Invoice Item ID"
// @Param request body invoiceitemdto.UpdateInvoiceItemRequest true "Update quantity"
//
// @Success 200 {object} docs.UpdateInvoiceItemSuccessResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
//
// @Router /invoice/{id}/items/{itemId} [patch]
func (c *InvoiceItemController) Update(
	w http.ResponseWriter,
	r *http.Request,
) error {

	var err error

	params := mux.Vars(r)

	invoiceID, _ := strconv.Atoi(params["id"])
	itemID, _ := strconv.Atoi(params["itemId"])

	req := invoiceitemdto.UpdateInvoiceItemRequest{}

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"json inválido",
		)
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			msg,
		)
	}

	res, err := c.service.Update(
		r.Context(),
		invoiceID,
		itemID,
		&req,
	)

	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(
		response.Success(res),
	)

	return nil
}

// Delete godoc
//
// @Summary Delete invoice item
// @Description Removes an item from a draft invoice
// @Tags Invoice Items
// @Produce json
// @Security BearerAuth
//
// @Param id path int true "Invoice ID"
// @Param itemId path int true "Invoice Item ID"
//
// @Success 200 {object} docs.DeleteInvoiceItemSuccessResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
//
// @Router /invoice/{id}/items/{itemId} [delete]
func (c *InvoiceItemController) Delete(
	w http.ResponseWriter,
	r *http.Request,
) error {

	var err error

	params := mux.Vars(r)

	invoiceID, _ := strconv.Atoi(params["id"])
	itemID, _ := strconv.Atoi(params["itemId"])

	err = c.service.Delete(
		r.Context(),
		invoiceID,
		itemID,
	)

	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(
		response.Success(
			"item eliminado correctamente",
		),
	)

	return nil
}

func (c *InvoiceItemController) GetInvoiceItems(
	w http.ResponseWriter,
	r *http.Request,
) error {
	var err error

	params := mux.Vars(r)

	invoiceID, err := strconv.Atoi(params["id"])

	if err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"invoice_id inválido",
		)
	}

	req, err := helper.ParseGetInvoicesRequest(r)
	if err != nil {
		return err
	}

	res, err := c.service.GetInvoiceItems(
		r.Context(),
		invoiceID,
		req,
	)

	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(
		response.Success(res),
	)
	return nil
}
