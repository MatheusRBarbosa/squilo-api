package interfaces

import "github.com/matheusrbarbosa/gofin/domain/models"

type TransactionService interface {
	PrepareTransaction(vault models.Vault, transaction *models.Transaction) error
}
