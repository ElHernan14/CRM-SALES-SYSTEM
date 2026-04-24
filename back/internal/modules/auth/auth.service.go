package auth

import (
	errorHandler "crm-system-sales/internal/core"
	models "crm-system-sales/internal/models/auth"
	"crm-system-sales/internal/modules/users"
	"crm-system-sales/internal/utils"
)

type AuthService struct {
	UserRepo *users.UserRepository
}

func NewAuthService(userRepo *users.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Login(email, password string) (*models.UserLogin, error) {
	user, err := s.UserRepo.GetUserLogin(email)
	if err != nil {
		return nil, errorHandler.ErrUnauthorized
	}

	err = utils.CheckPassword(password, user.PasswordHash)
	if err != nil {
		return nil, errorHandler.ErrUnauthorized
	}

	return user, nil
}
