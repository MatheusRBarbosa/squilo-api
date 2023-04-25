package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	crossCutting "github.com/matheusrbarbosa/squilo/crosscutting"
	l "github.com/matheusrbarbosa/squilo/crosscutting/logger"
	"github.com/matheusrbarbosa/squilo/domain/interfaces"
	"github.com/matheusrbarbosa/squilo/domain/models"
)

var authedUser *models.User

type authService struct {
	secret string
	issure string
}

func AuthService() interfaces.AuthService {
	return &authService{
		secret: crossCutting.GetEnv("JWT_SECRET"),
		issure: "squilo",
	}
}

func (s *authService) Generate(user models.User) string {
	claims := &models.UserCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7 * 2)),
			Issuer:    s.issure,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		l.GetLogger().Panicln(err)
	}

	return jwtToken
}

func (s *authService) Validate(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &models.UserCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])

		}
		return []byte(s.secret), nil
	})
}

func (s *authService) SetAuthUser(user models.User) error {
	authedUser = &user
	return nil
}

func (s *authService) GetAuthUser() *models.User {
	return authedUser
}
