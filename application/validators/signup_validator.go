package validators

import (
	"log"
	"time"

	"github.com/matheusrbarbosa/gofin/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	BirthDate string `json:"birthDate" binding:"required"` //TODO: Validar formato de data
	Password  string `json:"password" binding:"required,gte=6"`
}

func (r *SignupRequest) ParseToUser() models.User {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	birthDate, err := time.Parse("2006-01-02", r.BirthDate)
	if err != nil {
		log.Fatal(err)
	}

	return models.User{
		Name:      r.Name,
		Email:     r.Email,
		Password:  string(encryptedPassword),
		BirthDate: birthDate,
	}
}
