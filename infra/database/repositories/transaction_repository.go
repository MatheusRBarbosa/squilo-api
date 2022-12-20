package repositories

import (
	"github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/domain/models"
	"github.com/matheusrbarbosa/gofin/infra/database"
	"gorm.io/gorm"
)

type transactionRepository struct {
	context *gorm.DB
}

func TransactionRepository() interfaces.TransactionRepository {
	return &transactionRepository{
		context: database.Context(),
	}
}

func (r *transactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
	result := r.context.Create(&transaction)
	if result.Error != nil {
		return transaction, result.Error
	}

	return transaction, nil
}
