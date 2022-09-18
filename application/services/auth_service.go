package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	crossCutting "github.com/matheusrbarbosa/gofin/crosscutting"
	"github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/domain/models"
)

type jwtServices struct {
	secret string
	issure string
}

func JWTService() interfaces.JWTService {
	return &jwtServices{
		secret: crossCutting.AppEnvs.JWT_SECRET,
		issure: "matheusrbarbosa",
	}
}

func (service *jwtServices) Generate(user models.User) string {
	claims := &models.UserCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(service.secret))
	if err != nil {
		panic(err)
	}

	return jwtToken
}

func (service *jwtServices) Validate(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secret), nil
	})
}
