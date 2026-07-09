package invoiceController

import (
	errorHandler "crm-system-sales/internal/core/error"
	helper "crm-system-sales/internal/core/helper"
	"crm-system-sales/internal/core/response"
	validatorx "crm-system-sales/internal/core/validator"
	invoicedto "crm-system-sales/internal/modules/invoice/dto"
	invoiceService "crm-system-sales/internal/modules/invoice/service"
	submitInvoiceWorkflow "crm-system-sales/internal/services/invoice_workflow/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InvoiceController struct {
	Service        invoiceService.InvoiceService
	SubmitWorkflow submitInvoiceWorkflow.SubmitInvoiceWorkflow
}

func NewInvoiceController(
	service invoiceService.InvoiceService,
	submitWorkflow submitInvoiceWorkflow.SubmitInvoiceWorkflow,
) *InvoiceController {
	return &InvoiceController{
		Service:        service,
		SubmitWorkflow: submitWorkflow,
	}
}

// CreateDraft godoc
//
// @Summary Create invoice draft
// @Description Creates a new draft invoice for a buyer and seller relationship
// @Tags Invoice
// @Accept json
// @Produce json
// @Security BearerAuth
//
// @Param request body invoicedto.CreateInvoiceRequest true "Invoice data"
//
// @Success 200 {object} docs.CreateInvoiceSuccessResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
//
// @Router /invoice [post]
func (c *InvoiceController) CreateDraft(
	w http.ResponseWriter,
	r *http.Request,
) error {

	var req invoicedto.CreateInvoiceRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"JSON inválido",
		)
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			msg,
		)
	}

	res, err := c.Service.CreateDraft(r.Context(), &req)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(
		response.Success(res),
	)
}

// Submit godoc
//
// @Summary Submit invoice
// @Description Validates invoice items, reserves stock and changes invoice status from draft to pending
// @Tags Invoice
// @Accept json
// @Produce json
// @Security BearerAuth
//
// @Param id path int true "Invoice ID"
//
// @Success 200 {object} docs.SubmitInvoiceSuccessResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
//
// @Router /invoice/{id}/submit [post]
func (c *InvoiceController) Submit(
	w http.ResponseWriter,
	r *http.Request,
) error {
	var err error

	params := mux.Vars(r)

	if _, ok := params["id"]; !ok {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"ID de factura es requerido",
		)
	}

	invoiceID, err := strconv.Atoi(params["id"])
	if err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"ID de factura inválido",
		)
	}

	err = c.SubmitWorkflow.Submit(
		r.Context(),
		invoiceID,
	)

	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(
		response.Success(
			"invoice enviada correctamente",
		),
	)
}

// Pay godoc
//
// @Summary Pay invoice
// @Description Registers a payment for an invoice and updates payment status
// @Tags Invoice
// @Accept json
// @Produce json
// @Security BearerAuth
//
// @Param id path int true "Invoice ID"
// @Param request body invoicedto.PayInvoiceRequest true "Payment data"
//
// @Success 200 {object} docs.PayInvoiceSuccessResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
//
// @Router /invoice/{id}/pay [post]
func (c *InvoiceController) Pay(
	w http.ResponseWriter,
	r *http.Request,
) error {

	var err error

	params := mux.Vars(r)

	invoiceID, err := strconv.Atoi(params["id"])
	if err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"id inválido",
		)
	}

	var req invoicedto.PayInvoiceRequest

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"body inválido",
		)
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			msg,
		)
	}

	res, err := c.Service.Pay(
		r.Context(),
		invoiceID,
		&req,
	)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(
		response.Success(res),
	)
}

// GetByID godoc
//
// @Summary Get invoice by ID
// @Description Returns invoice details and financial information
// @Tags Invoice
// @Produce json
// @Security BearerAuth
//
// @Param id path int true "Invoice ID"
//
// @Success 200 {object} docs.GetInvoiceSuccessResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
//
// @Router /invoice/{id} [get]
func (c *InvoiceController) GetByID(
	w http.ResponseWriter,
	r *http.Request,
) error {

	id, err := strconv.Atoi(
		mux.Vars(r)["id"],
	)
	if err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"id inválido",
		)
	}

	res, err := c.Service.GetByID(
		r.Context(),
		id,
	)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).
		Encode(response.Success(res))
}

func (c *InvoiceController) Cancel(
	w http.ResponseWriter,
	r *http.Request,
) error {

	id, err := strconv.Atoi(
		mux.Vars(r)["id"],
	)

	if err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"id inválido",
		)
	}

	err = c.Service.Cancel(
		r.Context(),
		id,
	)

	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(
		response.Success(
			"Invoice cancelada correctamente",
		),
	)
}

func (c *InvoiceController) GetCompanyInvoices(
	w http.ResponseWriter,
	r *http.Request,
) error {

	req, err := helper.ParseGetCompanyInvoicesRequest(r)
	if err != nil {
		return err
	}

	res, err := c.Service.GetCompanyInvoices(
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
func (c *InvoiceController) GetCompanyPurchases(
	w http.ResponseWriter,
	r *http.Request,
) error {
	req, err := helper.ParseGetCompanyInvoicesRequest(r)
	if err != nil {
		return err
	}

	res, err := c.Service.GetCompanyPurchases(r.Context(), req)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}
