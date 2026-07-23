package authdto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type PersonalRegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,strong_password"`
	Phone     string `json:"phone" validate:"omitempty,max=50"`
}

type BusinessCompanyRegisterRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=150"`
	Email       string  `json:"email" validate:"required,email"`
	CategoryID  int     `json:"category_id" validate:"required,gt=0"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
}

type BusinessRegisterRequest struct {
	FirstName string                         `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string                         `json:"last_name" validate:"required,min=2,max=100"`
	Email     string                         `json:"email" validate:"required,email"`
	Password  string                         `json:"password" validate:"required,strong_password"`
	Phone     string                         `json:"phone" validate:"omitempty,max=50"`
	Company   BusinessCompanyRegisterRequest `json:"company" validate:"required"`
}

type AuthUserContextResponse struct {
	UserID      int      `json:"userID"`
	Email       string   `json:"email"`
	CompanyID   *int     `json:"company_id"`
	ClientID    *int     `json:"client_id"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type RegisterResponse struct {
	Token string                  `json:"token"`
	User  AuthUserContextResponse `json:"user"`
}
