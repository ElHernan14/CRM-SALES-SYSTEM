package invoiceitem

import (
	errorHandler "crm-system-sales/internal/core/error"
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

	return json.NewEncoder(w).Encode(
		response.Success(res),
	)
}
