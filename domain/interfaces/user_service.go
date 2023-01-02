package interfaces

import "github.com/matheusrbarbosa/squilo/domain/models"

type UserService interface {
	PrepareToCreate(user *models.User)
}
