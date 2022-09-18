package repository

import (
	"github.com/matheusrbarbosa/gofin/domain/models"
	"github.com/matheusrbarbosa/gofin/infra/database"
)

func Create(user models.User) models.User {
	result := database.DbContext.Create(&user)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	return user
}
