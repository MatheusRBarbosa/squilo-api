package services

import (
	"github.com/matheusrbarbosa/squilo/domain/exceptions"
	"github.com/matheusrbarbosa/squilo/domain/interfaces"
	"github.com/matheusrbarbosa/squilo/domain/models"
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

	return nil
}
