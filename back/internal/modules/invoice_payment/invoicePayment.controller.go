package invoicepayment

import (
	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/helper"
	"crm-system-sales/internal/core/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InvoicePaymentController struct {
	service InvoicePaymentService
}

func NewInvoicePaymentController(service InvoicePaymentService) *InvoicePaymentController {
	return &InvoicePaymentController{service: service}
}

func (c *InvoicePaymentController) GetByInvoice(
	w http.ResponseWriter,
	r *http.Request,
) error {
	params := mux.Vars(r)

	invoiceID, err := strconv.Atoi(params["id"])

	if err != nil {
		return errorHandler.NewAppError(
			http.StatusBadRequest,
			"invoice_id inválido",
		)
	}

	req, err := helper.ParseGetInvoicePaymentsRequest(r)
	if err != nil {
		return err
	}

	res, err := c.service.GetByInvoice(
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
