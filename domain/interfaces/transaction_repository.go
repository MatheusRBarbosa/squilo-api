package interfaces

import "github.com/matheusrbarbosa/gofin/domain/models"

type TransactionRepository interface {
	Create(models.Transaction) models.Transaction
	// GetById(id int) (models.Transaction, error)
}
