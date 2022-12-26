package repositories

import (
	"github.com/matheusrbarbosa/gofin/domain/interfaces"
	m "github.com/matheusrbarbosa/gofin/domain/models"
	"github.com/matheusrbarbosa/gofin/domain/utils"
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

func (r *vaultRepository) GetById(id int) (m.Vault, error) {
	vault := m.Vault{}
	err := r.context.Preload("Type").Preload("Transactions").First(&vault, id).Error
	if err == gorm.ErrRecordNotFound {
		return vault, gorm.ErrRecordNotFound
	}

	return vault, err
}

func (r *vaultRepository) GetByIdWithIncludes(id int) (m.Vault, error) {
	vault := m.Vault{}
	err := r.context.Preload(clause.Associations).First(&vault, id).Error
	if err == gorm.ErrRecordNotFound {
		return vault, gorm.ErrRecordNotFound
	}

	return vault, err
}

func (r *vaultRepository) Save(vault m.Vault) error {
	return r.context.Save(&vault).Error
}

func (r *vaultRepository) GetByUserId(userId int, pagination utils.Pagination) ([]m.Vault, error) {
	vaults := []m.Vault{}
	result := r.context.Where(&m.Vault{UserId: userId}).Offset(pagination.Offset).Limit(pagination.Limit).Find(&vaults)
	if result.Error != nil {
		return []m.Vault{}, result.Error
	}

	return vaults, nil
}
