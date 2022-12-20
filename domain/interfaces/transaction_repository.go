package interfaces

import "github.com/matheusrbarbosa/gofin/domain/models"

type TransactionRepository interface {
	Create(models.Transaction) (models.Transaction, error)
	// GetById(id int) (models.Transaction, error)
}
