package interfaces

import "github.com/matheusrbarbosa/gofin/domain/models"

type UserService interface {
	PrepareToCreate(user *models.User)
}
