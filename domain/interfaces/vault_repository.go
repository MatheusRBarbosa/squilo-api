package interfaces

import (
	"github.com/matheusrbarbosa/gofin/domain/models"
	"github.com/matheusrbarbosa/gofin/domain/utils"
)

type VaultRepository interface {
	GetByUserId(int, utils.Pagination) ([]models.Vault, error)
	GetById(id int) (models.Vault, error)
	GetByIdWithIncludes(id int) (models.Vault, error)
	Save(models.Vault) error
}
