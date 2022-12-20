package repositories

import (
	"github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/domain/models"
	"github.com/matheusrbarbosa/gofin/infra/database"
	"gorm.io/gorm"
)

type vaultRepository struct {
	context *gorm.DB
}

func VaultRepository() interfaces.VaultRepository {
	return &vaultRepository{
		context: database.Context(),
	}
}

func (r *vaultRepository) GetById(id int) (models.Vault, error) {
	vault := models.Vault{}
	err := r.context.Preload("Type").First(&vault, id).Error
	if err == gorm.ErrRecordNotFound {
		return vault, gorm.ErrRecordNotFound
	}

	return vault, err
}
