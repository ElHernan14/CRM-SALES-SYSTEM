package auth

import (
	"encoding/json"
	"log"
	"net/http"

	_ "crm-system-sales/docs"
	authcore "crm-system-sales/internal/core/auth"
	errorHandler "crm-system-sales/internal/core/error"
	response "crm-system-sales/internal/core/response"
	validatorx "crm-system-sales/internal/core/validator"
	authdto "crm-system-sales/internal/modules/auth/dto"
)

type AuthController struct {
	AuthService AuthService
}

func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

// Login godoc
//
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
//
// @Param request body authdto.LoginRequest true "Credentials"
//
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
//
// @Router /auth/login [post]
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) error {
	var req authdto.LoginRequest
	var err error

	err = json.NewDecoder(r.Body).Decode(&req)
	log.Println("Received login request: ", req)
	if err != nil {
		log.Println("Error decoding login request: ", err)
		err = errorHandler.NewAppError(http.StatusBadRequest, "Invalid request payload")
		return err
	}

	msg, validate := validatorx.ValidateStruct(req)
	if validate {
		log.Printf("Invalid Request: %s", msg)
		err = errorHandler.NewAppError(http.StatusBadRequest, msg)
		return err
	}

	user, err := c.AuthService.Login(r, req.Email, req.Password)
	if err != nil {
		log.Println("Error logging in user: ", err)
		err = errorHandler.NewAppError(http.StatusUnauthorized, "invalid credentials")
		return err
	}

	token, err := authcore.GenerateJWT(user, user.Permissions, user.Roles)
	if err != nil {
		log.Println("Error generating JWT: ", err)
		err = errorHandler.NewAppError(http.StatusInternalServerError, "internal server error")
		return err
	}

	res := &authdto.LoginResponse{
		Token: token,
	}

	log.Print("Login success")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.Success(res))
	return nil
}
