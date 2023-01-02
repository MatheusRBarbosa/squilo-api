package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/squilo/application/services"
	"github.com/matheusrbarbosa/squilo/application/validators"
	"github.com/matheusrbarbosa/squilo/domain/exceptions"
	i "github.com/matheusrbarbosa/squilo/domain/interfaces"
	"github.com/matheusrbarbosa/squilo/infra/database/repositories"
	"golang.org/x/crypto/bcrypt"
)

type authHandler struct {
	userRepository i.UserRepository
	authService    i.AuthService
}

func AuthHandler() *authHandler {
	return &authHandler{
		userRepository: repositories.UserRepository(),
		authService:    services.AuthService(),
	}
}

func (h *authHandler) HandleLogin(request validators.LoginRequest) (*gin.H, error) {
	user, err := h.userRepository.GetByEmail(request.Email)
	invalidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil || invalidPassword != nil {
		return nil, exceptions.UNAUTHORIZED
	}

	token := h.authService.Generate(user)
	return &gin.H{"token": token}, nil
}
