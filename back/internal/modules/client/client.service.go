package client

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	constants "crm-system-sales/internal/constants"
	core "crm-system-sales/internal/core"
	"crm-system-sales/internal/core/access"
	errorHandler "crm-system-sales/internal/core/error"
	"crm-system-sales/internal/core/utils"
	validatorx "crm-system-sales/internal/core/validator"

	clientdto "crm-system-sales/internal/dto"
	clientModel "crm-system-sales/internal/models/client"

	userModel "crm-system-sales/internal/models/users"
	"crm-system-sales/internal/modules/auth"
	"crm-system-sales/internal/modules/users"

	"golang.org/x/crypto/bcrypt"
)

type ClientService interface {
	Create(ctx context.Context, req *clientdto.CreateClientRequest) (*clientdto.ClientResponse, error)
	GetClients(ctx context.Context, req *clientdto.GetClientsRequest) (*clientdto.GetClientsResponse, error)
}

type clientService struct {
	db       *sql.DB
	repo     ClientRepository
	userRepo users.UserRepository
	authRepo auth.AuthRepository
}

func NewClientService(db *sql.DB, repo ClientRepository, userRepo users.UserRepository, authRepo auth.AuthRepository) ClientService {
	return &clientService{
		db:       db,
		repo:     repo,
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

func (s *clientService) Create(ctx context.Context, req *clientdto.CreateClientRequest) (*clientdto.ClientResponse, error) {
	// Tracing
	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE CreateClient")(err)
	}()

	//  validar DTO
	msg, invalid := validatorx.ValidateStruct(req)
	if invalid {
		err = errorHandler.NewAppError(400, msg)
		return nil, err
	}

	//  validar email
	exists, err := s.repo.EmailExists(req.Email)
	if err != nil {
		log.Println("Error checking client email:", err)
		return nil, err
	}
	if exists {
		log.Println("Email client already exists:", req.Email)
		err = errorHandler.NewAppError(http.StatusConflict, "Email client already exists")
		return nil, err
	}

	//  contexto multi-tenant
	tenant := core.GetTenant(ctx)
	var companyID *int
	if tenant != nil && tenant.CompanyID != nil {
		companyID = tenant.CompanyID
	}
	if req.CompanyID != nil {
		companyID = req.CompanyID
	}

	var clientResp *clientdto.ClientResponse

	//  transacción centralizada
	err = core.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {
		// crear user
		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user := &userModel.User{
			Email:        req.Email,
			PasswordHash: string(hash),
			CompanyID:    companyID,
		}

		exists, err := s.userRepo.EmailExists(req.Email)
		if err != nil {
			log.Println("Error checking user email:", err)
			return err
		}
		if exists {
			log.Println("Email user already exists:", req.Email)
			err = errorHandler.NewAppError(http.StatusConflict, "Email user already exists")
			return err
		}

		user, err = s.userRepo.Create(tx, user)
		if err != nil {
			log.Println("Error creating user:", err)
			err = errorHandler.NewAppError(http.StatusInternalServerError, "Failed to create user")
			return err
		}

		// asignar rol
		roleName := constants.RoleManager
		if companyID == nil {
			roleName = constants.RoleAdmin
		}
		if err := s.authRepo.AssignRole(tx, user.ID, roleName); err != nil {
			log.Println("Error assigning role:", err)
			err = errorHandler.NewAppError(http.StatusInternalServerError, "Failed to assign role")
			return err
		}

		// crear client
		client := &clientModel.Client{
			UserID:    &user.ID,
			CompanyID: companyID,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
		}

		if err := s.repo.Create(tx, client); err != nil {
			log.Println("Error creating client:", err)
			err = errorHandler.NewAppError(http.StatusInternalServerError, "Failed to create client")
			return err
		}

		log.Println("Client created successfully:", client.ID)

		clientResp = &clientdto.ClientResponse{
			ID:        client.ID,
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Email:     client.Email,
			CompanyID: client.CompanyID,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return clientResp, nil
}

func (s *clientService) GetClients(ctx context.Context, req *clientdto.GetClientsRequest) (*clientdto.GetClientsResponse, error) {

	var err error
	defer func() {
		utils.Trace(ctx, "SERVICE GetClients")(err)
	}()

	tenant := core.GetTenant(ctx)

	companyID, err := access.ResolveClientScope(tenant, req.CompanyID)
	if err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.Limit

	clients, total, err := s.repo.GetClients(
		ctx,
		companyID,
		req.Search,
		req.Email,
		req.Limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	resp := &clientdto.GetClientsResponse{
		Clients: []clientdto.ClientResponse{},
		Meta: clientdto.Meta{
			Page:  req.Page,
			Limit: req.Limit,
			Total: total,
		},
	}

	// calcular total pages
	if total > 0 {
		resp.Meta.TotalPages = (total + req.Limit - 1) / req.Limit
	}

	for _, c := range clients {
		resp.Clients = append(resp.Clients, clientdto.ClientResponse{
			ID:        c.ID,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Email:     c.Email,
			CompanyID: c.CompanyID,
		})
	}

	return resp, nil
}
