package repositories

import (
	"github.com/matheusrbarbosa/squilo/domain/interfaces"
	m "github.com/matheusrbarbosa/squilo/domain/models"
	"github.com/matheusrbarbosa/squilo/domain/utils"
	"github.com/matheusrbarbosa/squilo/infra/database"
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

func (r *transactionRepository) Create(transaction m.Transaction) (m.Transaction, error) {
	result := r.context.Create(&transaction)
	if result.Error != nil {
		return transaction, result.Error
	}

	return transaction, nil
}

func (r *transactionRepository) GetById(id int) (m.Transaction, error) {
	transaction := m.Transaction{}
	err := r.context.Preload("Vault").First(&transaction, id).Error
	if err == gorm.ErrRecordNotFound {
		return transaction, gorm.ErrRecordNotFound
	}

	return transaction, err
}

func (r *transactionRepository) Delete(id int) error {
	return r.context.Delete(&m.Transaction{}, id).Error
}

func (r *transactionRepository) Update(transaction *m.Transaction, newTransaction m.Transaction) error {
	err := r.context.Model(&transaction).Updates(m.Transaction{
		Value:       newTransaction.Value,
		Date:        newTransaction.Date,
		Observation: newTransaction.Observation,
	}).Error

	return err
}

func (r *transactionRepository) GetByVaultId(id int, pagination utils.Pagination) ([]m.Transaction, error) {
	transactions := []m.Transaction{}
	result := r.context.Where(&m.Transaction{VaultId: id}).Offset(pagination.Offset).Limit(pagination.Limit).Find(&transactions)
	if result.Error != nil {
		return []m.Transaction{}, result.Error
	}

	return transactions, nil
}
