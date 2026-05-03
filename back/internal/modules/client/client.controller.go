package client

import (
	errorHandler "crm-system-sales/internal/core/error"
	helper "crm-system-sales/internal/core/helper"
	response "crm-system-sales/internal/core/response"
	"crm-system-sales/internal/core/utils"
	validatorx "crm-system-sales/internal/core/validator"
	clientdto "crm-system-sales/internal/dto"
	"encoding/json"
	"net/http"
)

type ClientController struct {
	service ClientService
}

func NewClientController(service ClientService) *ClientController {
	return &ClientController{service: service}
}

func (c *ClientController) Create(w http.ResponseWriter, r *http.Request) error {
	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER CreateClient")(err)
	}()

	var req clientdto.CreateClientRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errorHandler.NewAppError(http.StatusBadRequest, "Invalid request payload")
	}

	res, err := c.service.Create(r.Context(), &req)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(response.Success(res))
	return nil
}

func (c *ClientController) GetClients(w http.ResponseWriter, r *http.Request) error {

	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER GetClients")(err)
	}()

	req, err := helper.ParseGetClientsRequest(r)
	if err != nil {
		return err
	}

	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		return errorHandler.NewAppError(http.StatusBadRequest, msg)
	}

	res, err := c.service.GetClients(r.Context(), req)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(response.Success(res))
	return nil
}
