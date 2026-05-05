package client

import (
	errorHandler "crm-system-sales/internal/core/error"
	helper "crm-system-sales/internal/core/helper"
	response "crm-system-sales/internal/core/response"
	utils "crm-system-sales/internal/core/utils"
	validatorx "crm-system-sales/internal/core/validator"
	clientdto "crm-system-sales/internal/modules/client/dto"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (c *ClientController) GetClientByID(w http.ResponseWriter, r *http.Request) error {

	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER GetClientByID")(err)
	}()

	vars := mux.Vars(r)

	idStr := vars["id"]
	id, convErr := strconv.Atoi(idStr)
	if convErr != nil || id <= 0 {
		err = errorHandler.NewAppError(http.StatusBadRequest, "id inválido ó no proporcionado")
		return err
	}

	res, err := c.service.GetClientByID(r.Context(), id)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(response.Success(res))
	return nil
}

func (c *ClientController) UpdateClient(w http.ResponseWriter, r *http.Request) error {

	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER UpdateClient")(err)
	}()

	vars := mux.Vars(r)

	id, convErr := strconv.Atoi(vars["id"])
	if convErr != nil || id <= 0 {
		err = errorHandler.NewAppError(http.StatusBadRequest, "id inválido ó no proporcionado")
		return err
	}

	var req clientdto.UpdateClientRequest

	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		err = errorHandler.NewAppError(http.StatusBadRequest, "body inválido")
		return err
	}

	// validar DTO
	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		err = errorHandler.NewAppError(http.StatusBadRequest, msg)
		return err
	}

	res, err := c.service.UpdateClient(r.Context(), id, &req)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(response.Success(res))
	return nil
}

func (c *ClientController) DeleteClient(w http.ResponseWriter, r *http.Request) error {

	var err error
	defer func() {
		utils.Trace(r.Context(), "CONTROLLER DeleteClient")(err)
	}()

	vars := mux.Vars(r)

	id, convErr := strconv.Atoi(vars["id"])
	if convErr != nil || id <= 0 {
		err = errorHandler.NewAppError(http.StatusBadRequest, "id inválido ó no proporcionado")
		return err
	}

	err = c.service.DeleteClient(r.Context(), id)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(response.Success(nil))
	return nil
}
