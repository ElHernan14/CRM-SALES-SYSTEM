package invoice

import (
	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/response"
	validatorx "crm-system-sales/internal/core/validator"
	invoicedto "crm-system-sales/internal/modules/invoice/dto"
	"encoding/json"
	"net/http"
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
