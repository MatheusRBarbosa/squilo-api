package services

import (
	"time"

	"github.com/matheusrbarbosa/gofin/domain/exceptions"
	"github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/domain/models"
)

type transactionService struct{}

func TransactionService() interfaces.TransactionService {
	return &transactionService{}
}

func (s *transactionService) PrepareTransaction(v models.Vault, t *models.Transaction) error {
	negativeVaule := t.Value < 0

	if (negativeVaule && !v.Type.AllowNegativeTransactions) || (!negativeVaule && !v.Type.AllowPositiveTransactions) {
		return exceptions.TRANSACTION_WRONG_TYPE
	}

	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	return nil
}
