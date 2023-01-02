package interfaces

import (
	"github.com/matheusrbarbosa/squilo/domain/models"
	"github.com/matheusrbarbosa/squilo/domain/utils"
)

type TransactionRepository interface {
	Create(models.Transaction) (models.Transaction, error)
	GetById(id int) (models.Transaction, error)
	Delete(id int) error
	Update(*models.Transaction, models.Transaction) error
	GetByVaultId(int, utils.Pagination) ([]models.Transaction, error)
}
