package handlers

import (
	v "github.com/matheusrbarbosa/gofin/application/validators"
	"github.com/matheusrbarbosa/gofin/domain/dtos"
	"github.com/matheusrbarbosa/gofin/domain/exceptions"
	"github.com/matheusrbarbosa/gofin/infra/database/repositories"
	"gorm.io/gorm"
)

func HandleSignup(request v.SignupRequest) (dtos.UserDto, error) {
	user := request.ParseToUser()
	userRepository := repositories.UserRepository()

	_, err := userRepository.GetByEmail(user.Email)
	if err != nil && err == gorm.ErrRecordNotFound {
		user = userRepository.Create(user)
		return user.ParseDto(), nil
	} else {
		return user.ParseDto(), exceptions.EMAIL_ALREADY_EXIST
	}
}
