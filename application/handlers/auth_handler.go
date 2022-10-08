package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/application/services"
	"github.com/matheusrbarbosa/gofin/application/validators"
	"github.com/matheusrbarbosa/gofin/domain/exceptions"
	"github.com/matheusrbarbosa/gofin/infra/database/repositories"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(context *gin.Context, request validators.LoginRequest) (*gin.H, error) {
	user, err := repositories.UserRepository().GetByEmail(request.Email)
	invalidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil || invalidPassword != nil {
		return nil, exceptions.UNAUTHORIZED
	}

	token := services.JWTService().Generate(user)
	return &gin.H{"token": token}, nil
}
