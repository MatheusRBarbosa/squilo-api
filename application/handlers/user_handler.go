package handlers

import (
	v "github.com/matheusrbarbosa/gofin/application/validators"
	"github.com/matheusrbarbosa/gofin/domain/dtos"
	"github.com/matheusrbarbosa/gofin/infra/database/repository"
)

func HandleSignup(request v.SignupRequest) dtos.UserDto {
	user := request.ParseToUser()

	user = repository.Create(user)

	return user.ParseDto()
}
