package interfaces

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/matheusrbarbosa/squilo/domain/models"
)

type AuthService interface {
	Generate(user models.User) string
	Validate(token string) (*jwt.Token, error)
	SetAuthUser(user models.User) error
	GetAuthUser() *models.User
}
