package repositories

import (
	"github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/domain/models"
	"github.com/matheusrbarbosa/gofin/infra/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	err := r.context.Preload("Type").Preload("Transactions").First(&vault, id).Error
	if err == gorm.ErrRecordNotFound {
		return vault, gorm.ErrRecordNotFound
	}

	return vault, err
}

func (r *vaultRepository) GetByIdWithIncludes(id int) (models.Vault, error) {
	vault := models.Vault{}
	err := r.context.Preload(clause.Associations).First(&vault, id).Error
	if err == gorm.ErrRecordNotFound {
		return vault, gorm.ErrRecordNotFound
	}

	return vault, err
}

func (r *vaultRepository) Save(vault models.Vault) error {
	return r.context.Save(&vault).Error
}
