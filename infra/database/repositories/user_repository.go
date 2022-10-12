package repositories

import (
	"github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/domain/models"
	"github.com/matheusrbarbosa/gofin/infra/database"
)

type userRepository struct{}

func UserRepository() interfaces.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Create(user models.User) models.User {
	result := database.DbContext.Create(&user)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	return user
}

func (ur *userRepository) GetByEmail(email string) (models.User, error) {
	user := models.User{}
	err := database.DbContext.Where(&models.User{Email: email}).First(&user).Error

	return user, err
}

func (ur *userRepository) GetById(id int) (models.User, error) {
	user := models.User{}
	err := database.DbContext.Where(&models.User{ID: id}).First(&user).Error

	return user, err
}
