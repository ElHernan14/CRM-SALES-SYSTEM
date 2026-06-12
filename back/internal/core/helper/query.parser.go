package helper

import (
	errorHandler "crm-system-sales/internal/core/error"
	validatorx "crm-system-sales/internal/core/validator"
	clientdto "crm-system-sales/internal/modules/client/dto"
	invoicedto "crm-system-sales/internal/modules/invoice/dto"
	invoiceitemdto "crm-system-sales/internal/modules/invoice_item/dto"
	invoicepaymentdto "crm-system-sales/internal/modules/invoice_payment/dto"
	productdto "crm-system-sales/internal/modules/product/dto"

	"net/http"
	"strconv"
)

func ParseGetClientsRequest(r *http.Request) (*clientdto.GetClientsRequest, error) {
	q := r.URL.Query()

	req := &clientdto.GetClientsRequest{}

	req.Search = q.Get("search")
	req.Email = q.Get("email")

	if v := q.Get("company_id"); v != "" {
		id, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "company_id inválido")
		}
		req.CompanyID = &id
	}

	if v := q.Get("page"); v != "" {
		page, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "page inválido")
		}
		req.Page = page
	}

	if v := q.Get("limit"); v != "" {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "limit inválido")
		}
		req.Limit = limit
	}

	// defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	// validate
	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	return req, nil
}

func ParseGetInvoicesRequest(r *http.Request) (*invoiceitemdto.GetInvoiceItemsRequest, error) {
	q := r.URL.Query()

	req := &invoiceitemdto.GetInvoiceItemsRequest{}

	if v := q.Get("page"); v != "" {
		page, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "page inválido")
		}
		req.Page = page
	}

	if v := q.Get("limit"); v != "" {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "limit inválido")
		}
		req.Limit = limit
	}

	// defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	// validate
	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	return req, nil
}

func ParseGetInvoicePaymentsRequest(r *http.Request) (*invoicepaymentdto.GetInvoicePaymentsRequest, error) {
	q := r.URL.Query()

	req := &invoicepaymentdto.GetInvoicePaymentsRequest{}
	req.PaymentMethod = q.Get("payment_method")
	req.SortColumn = q.Get("sort_column")
	req.Order = q.Get("order")

	if v := q.Get("page"); v != "" {
		page, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "page inválido")
		}
		req.Page = page
	}

	if v := q.Get("limit"); v != "" {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "limit inválido")
		}
		req.Limit = limit
	}

	// defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	// validate
	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	return req, nil
}

func ParseGetCompanyInvoicesRequest(r *http.Request) (*invoicedto.GetCompanyInvoicesRequest, error) {
	q := r.URL.Query()

	req := &invoicedto.GetCompanyInvoicesRequest{}
	req.StatusInvoice = q.Get("status_invoice")
	req.SortColumn = q.Get("sort_column")
	req.Order = q.Get("order")

	if v := q.Get("page"); v != "" {
		page, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "page inválido")
		}
		req.Page = page
	}

	if v := q.Get("limit"); v != "" {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "limit inválido")
		}
		req.Limit = limit
	}

	if v := q.Get("status"); v != "" {
		status, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "status inválido")
		}
		req.Status = status
	}

	// defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	// validate
	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	return req, nil
}

func ParseGetCompanyProductsRequest(r *http.Request) (*productdto.GetCompanyProductsRequest, error) {
	q := r.URL.Query()

	req := &productdto.GetCompanyProductsRequest{}
	req.Name = q.Get("name")
	req.Type = q.Get("type")
	req.SortColumn = q.Get("sort_column")
	req.Order = q.Get("order")

	if v := q.Get("page"); v != "" {
		page, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "page inválido")
		}
		req.Page = page
	}

	if v := q.Get("limit"); v != "" {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "limit inválido")
		}
		req.Limit = limit
	}

	if v := q.Get("status"); v != "" {
		status, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "status inválido")
		}
		req.Status = status
	}

	// defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	// validate
	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	return req, nil
}

func ParseGetCompanyCustomersRequest(r *http.Request) (*clientdto.GetCompanyCustomersRequest, error) {
	q := r.URL.Query()

	req := &clientdto.GetCompanyCustomersRequest{}
	req.Name = q.Get("name")
	req.SortColumn = q.Get("sort_column")
	req.Order = q.Get("order")

	if v := q.Get("page"); v != "" {
		page, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "page inválido")
		}
		req.Page = page
	}

	if v := q.Get("limit"); v != "" {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return nil, errorHandler.NewAppError(http.StatusBadRequest, "limit inválido")
		}
		req.Limit = limit
	}

	// defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	// validate
	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return nil, errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	return req, nil
}
