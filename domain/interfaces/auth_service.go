package interfaces

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/matheusrbarbosa/gofin/domain/models"
)

type JWTService interface {
	Generate(user models.User) string
	Validate(token string) (*jwt.Token, error)
}
