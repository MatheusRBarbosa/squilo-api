package interfaces

import "github.com/matheusrbarbosa/gofin/domain/models"

type UserRepository interface {
	Create(user models.User) models.User
	GetByEmail(email string) (models.User, error)
}
