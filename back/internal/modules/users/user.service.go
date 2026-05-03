package users

import (
	users "crm-system-sales/internal/models/users"
)

type UserService interface {
	GetUsers() ([]users.User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUsers() ([]users.User, error) {
	return s.repo.GetAll()
}
