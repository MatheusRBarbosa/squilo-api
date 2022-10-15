package repositories

import (
	"github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/domain/models"
	"github.com/matheusrbarbosa/gofin/infra/database"
	"gorm.io/gorm"
)

type userRepository struct {
	context *gorm.DB
}

func UserRepository() interfaces.UserRepository {
	return &userRepository{
		context: database.Context(),
	}
}

func (ur *userRepository) Create(user models.User) models.User {
	result := ur.context.Create(&user)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	return user
}

func (ur *userRepository) GetByEmail(email string) (models.User, error) {
	user := models.User{}
	err := ur.context.Where(&models.User{Email: email}).First(&user).Error

	return user, err
}

func (ur *userRepository) GetById(id int) (models.User, error) {
	user := models.User{}
	err := ur.context.Where(&models.User{ID: id}).First(&user).Error

	return user, err
}
