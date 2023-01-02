package interfaces

import (
	"github.com/matheusrbarbosa/squilo/domain/models"
	"github.com/matheusrbarbosa/squilo/domain/utils"
)

type VaultRepository interface {
	GetByUserId(int, utils.Pagination) ([]models.Vault, error)
	GetById(id int) (models.Vault, error)
	GetByIdWithIncludes(id int) (models.Vault, error)
	Save(models.Vault) error
}
