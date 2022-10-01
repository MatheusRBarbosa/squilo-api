package validators

import (
	"github.com/matheusrbarbosa/gofin/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6"`
}

func (r *SignupRequest) ParseToUser() models.User {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return models.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: string(encryptedPassword),
	}
}
