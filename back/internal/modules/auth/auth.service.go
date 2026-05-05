package auth

import (
	authcore "crm-system-sales/internal/core/auth"
	errorHandler "crm-system-sales/internal/core/error"
	coreUtils "crm-system-sales/internal/core/utils"
	models "crm-system-sales/internal/models/auth"
	"crm-system-sales/internal/modules/users"
	"log"
	"net/http"
)

type AuthService interface {
	Login(r *http.Request, email, password string) (*models.UserLogin, error)
}

type authService struct {
	UserRepo users.UserRepository
}

func NewAuthService(userRepo users.UserRepository) AuthService {
	return &authService{UserRepo: userRepo}
}

func (s *authService) Login(r *http.Request, email, password string) (*models.UserLogin, error) {
	var err error
	defer func() {
		coreUtils.Trace(r.Context(), "SERVICE Login")(err)
	}()
	user, err := s.UserRepo.GetUserLogin(email)
	if err != nil {
		log.Println("Error fetching user for login: ", err)
		err = errorHandler.NewAppError(http.StatusUnauthorized, "invalid credentials")
		return nil, err
	}

	err = authcore.CheckPassword(password, user.PasswordHash)
	if err != nil {
		log.Println("Error checking password: ", err)
		err = errorHandler.NewAppError(http.StatusUnauthorized, "invalid credentials")
		return nil, err
	}

	return user, nil
}
