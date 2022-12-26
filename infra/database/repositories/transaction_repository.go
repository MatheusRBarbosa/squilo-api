package repositories

import (
	"github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/domain/models"
	"github.com/matheusrbarbosa/gofin/domain/utils"
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

func (r *transactionRepository) GetById(id int) (models.Transaction, error) {
	transaction := models.Transaction{}
	err := r.context.Preload("Vault").First(&transaction, id).Error
	if err == gorm.ErrRecordNotFound {
		return transaction, gorm.ErrRecordNotFound
	}

	return transaction, err
}

func (r *transactionRepository) Delete(id int) error {
	return r.context.Delete(&models.Transaction{}, id).Error
}

func (r *transactionRepository) Update(transaction *models.Transaction, newTransaction models.Transaction) error {
	err := r.context.Model(&transaction).Updates(models.Transaction{
		Value:       newTransaction.Value,
		Date:        newTransaction.Date,
		Observation: newTransaction.Observation,
	}).Error

	return err
}

func (r *transactionRepository) GetByVaultId(id int, pagination utils.Pagination) ([]models.Transaction, error) {
	transactions := []models.Transaction{}
	result := r.context.Where(&models.Transaction{VaultId: id}).Offset(pagination.Offset).Limit(pagination.Limit).Find(&transactions)
	if result.Error != nil {
		return []models.Transaction{}, result.Error
	}

	return transactions, nil
}
