package invoice

import (
	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/response"
	"crm-system-sales/internal/core/utils"
	validatorx "crm-system-sales/internal/core/validator"
	invoicedto "crm-system-sales/internal/modules/invoice/dto"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InvoiceController struct {
	Service InvoiceService
}

func NewInvoiceController(service InvoiceService) *InvoiceController {
	return &InvoiceController{Service: service}
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

	invoiceID, _ := strconv.Atoi(params["id"])

	err = c.Service.Submit(
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
