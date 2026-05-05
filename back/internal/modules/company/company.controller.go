package company

import (
	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/response"
	"crm-system-sales/internal/core/utils"
	validatorx "crm-system-sales/internal/core/validator"
	companydto "crm-system-sales/internal/modules/company/dto"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CompanyController struct {
	service CompanyService
}

func NewCompanyController(service CompanyService) *CompanyController {
	return &CompanyController{service: service}
}

func (c *CompanyController) Create(w http.ResponseWriter, r *http.Request) error {

	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER CreateCompany")(err)
	}()

	var req companydto.CreateCompanyRequest

	if errDecode := json.NewDecoder(r.Body).Decode(&req); errDecode != nil {
		return errorHandler.NewAppError(http.StatusBadRequest, "body inválido")
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	res, err := c.service.CreateCompany(r.Context(), &req)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(response.Success(res))
	return nil
}

func (c *CompanyController) GetMyCompany(w http.ResponseWriter, r *http.Request) error {

	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER GetMyCompany")(err)
	}()

	res, err := c.service.GetMyCompany(r.Context())
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *CompanyController) GetCompanies(w http.ResponseWriter, r *http.Request) error {

	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER GetCompanies")(err)
	}()

	q := r.URL.Query()

	req := companydto.GetCompaniesRequest{
		Search: q.Get("search"),
	}

	req.Page = 1
	req.Limit = 10

	if v := q.Get("page"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "page inválido")
		}
		if p > 0 {
			req.Page = p
		}
	}

	if v := q.Get("limit"); v != "" {
		l, err := strconv.Atoi(v)
		if err != nil {
			return errorHandler.NewAppError(http.StatusBadRequest, "limit inválido")
		}
		if l > 0 {
			req.Limit = l
		}
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	res, err := c.service.GetCompanies(r.Context(), &req)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}

func (c *CompanyController) GetCompanyByID(w http.ResponseWriter, r *http.Request) error {

	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER GetCompanyByID")(err)
	}()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		return errorHandler.NewAppError(http.StatusBadRequest, "id inválido en path")
	}

	res, err := c.service.GetByID(r.Context(), id)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(response.Success(res))
}
