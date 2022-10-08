package handlers

import (
	v "github.com/matheusrbarbosa/gofin/application/validators"
	"github.com/matheusrbarbosa/gofin/domain/dtos"
	"github.com/matheusrbarbosa/gofin/infra/database/repositories"
)

func HandleSignup(request v.SignupRequest) dtos.UserDto {
	user := request.ParseToUser()

	user = repositories.UserRepository().Create(user)

	return user.ParseDto()
}
