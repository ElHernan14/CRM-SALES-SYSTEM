package authcore

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"

	models "crm-system-sales/internal/models/auth"
)

type Claims struct {
	UserID      int       `json:"user_id"`
	Email       string    `json:"email"`
	CompanyID   *int      `json:"company_id,omitempty"`
	Roles       *[]string `json:"roles,omitempty"`
	Permissions *[]string `json:"permissions"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("super_secret_key") // después va a .env

func GenerateJWT(user *models.UserLogin, permissions []string, roles []string) (string, error) {
	log.Println("User roles in GenerateJWT:", roles)
	log.Println("User permissions in GenerateJWT:", permissions)
	claims := Claims{
		UserID:      user.ID,
		Email:       user.Email,
		CompanyID:   user.CompanyID,
		Roles:       &roles,
		Permissions: &permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
