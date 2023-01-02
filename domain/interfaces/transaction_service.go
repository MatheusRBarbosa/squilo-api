package interfaces

import "github.com/matheusrbarbosa/squilo/domain/models"

type TransactionService interface {
	PrepareTransaction(vault models.Vault, transaction *models.Transaction) error
}
