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

func (r *userRepository) Create(user models.User) models.User {
	result := r.context.Create(&user)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	return user
}

func (r *userRepository) GetByEmail(email string) (models.User, error) {
	user := models.User{}
	err := r.context.Where(&models.User{Email: email}).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, gorm.ErrRecordNotFound
	}

	return user, err
}

func (r *userRepository) GetById(id int) (models.User, error) {
	user := models.User{}
	err := r.context.Preload("Vaults").Where(&models.User{ID: id}).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, gorm.ErrRecordNotFound
	}

	return user, err
}
