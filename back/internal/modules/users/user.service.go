package users

type UserService struct {
	UserRepository *UserRepository
}

func NewUserService(userRepo *UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}
