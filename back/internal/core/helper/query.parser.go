package helper

import (
	errorHandler "crm-system-sales/internal/core/error"
	validatorx "crm-system-sales/internal/core/validator"
	clientdto "crm-system-sales/internal/modules/client/dto"

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
