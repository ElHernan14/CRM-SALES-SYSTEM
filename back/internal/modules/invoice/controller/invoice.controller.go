package invoiceController

import (
	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/response"
	"crm-system-sales/internal/core/utils"
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

func (c *InvoiceController) Submit(
	w http.ResponseWriter,
	r *http.Request,
) error {
	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER SubmitInvoice")(err)
	}()

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

func (c *InvoiceController) Pay(
	w http.ResponseWriter,
	r *http.Request,
) error {

	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER PayInvoice")(err)
	}()

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
