package handlers

import (
	"github.com/matheusrbarbosa/squilo/application/services"
	v "github.com/matheusrbarbosa/squilo/application/validators"
	"github.com/matheusrbarbosa/squilo/crosscutting/logger"
	"github.com/matheusrbarbosa/squilo/domain/dtos"
	"github.com/matheusrbarbosa/squilo/domain/exceptions"
	i "github.com/matheusrbarbosa/squilo/domain/interfaces"
	"github.com/matheusrbarbosa/squilo/domain/models"
	"github.com/matheusrbarbosa/squilo/domain/utils"
	"github.com/matheusrbarbosa/squilo/infra/database"
	"github.com/matheusrbarbosa/squilo/infra/database/repositories"
	"gorm.io/gorm"
)

type transactionHandler struct {
	logger                logger.Logger
	db                    *gorm.DB
	transactionService    i.TransactionService
	vaultRepository       i.VaultRepository
	transactionRepository i.TransactionRepository
}

func TransactionHandler() *transactionHandler {
	return &transactionHandler{
		logger:                logger.GetLogger(),
		db:                    database.Context(),
		transactionService:    services.TransactionService(),
		vaultRepository:       repositories.VaultRepository(),
		transactionRepository: repositories.TransactionRepository(),
	}
}

func (h *transactionHandler) Create(vaultId int, request v.CreateTransactionRequest) (dtos.TransactionDto, error) {
	vault, err := h.vaultRepository.GetById(vaultId)
	if err != nil {
		h.logger.Errorf(err.Error())
		return dtos.TransactionDto{}, exceptions.VAULT_NOT_FOUND
	}

	transaction := request.ParseToTransaction(vaultId)
	if err = h.transactionService.PrepareTransaction(vault, &transaction); err != nil {
		h.logger.Errorf(err.Error())
		return dtos.TransactionDto{}, err
	}

	h.db.Transaction(func(tx *gorm.DB) error {
		transaction, err = h.transactionRepository.Create(transaction)
		if err != nil {
			h.logger.Errorf(err.Error())
			return err
		}

		vault.Total += transaction.Value
		err = h.vaultRepository.Save(vault)
		if err != nil {
			h.logger.Errorf(err.Error())
			return err
		}

		return nil
	})

	return transaction.ParseDto(), nil
}

func (h *transactionHandler) Delete(vaultId, transactionId int) (dtos.TransactionDto, error) {
	vault, transaction, err := h.getVaultAndTransaction(vaultId, transactionId)
	if err != nil {
		h.logger.Errorf(err.Error())
		return dtos.TransactionDto{}, err
	}

	h.db.Transaction(func(tx *gorm.DB) error {
		vault.Total -= transaction.Value
		err = h.vaultRepository.Save(vault)
		if err != nil {
			h.logger.Errorf(err.Error())
			return err
		}

		err = h.transactionRepository.Delete(transactionId)
		if err != nil {
			h.logger.Errorf(err.Error())
			return err
		}

		return nil
	})

	return transaction.ParseDto(), nil
}

func (h *transactionHandler) Update(vaultId int, transactionId int, request v.CreateTransactionRequest) (dtos.TransactionDto, error) {
	vault, transaction, err := h.getVaultAndTransaction(vaultId, transactionId)
	if err != nil {
		h.logger.Errorf(err.Error())
		return dtos.TransactionDto{}, err
	}

	newTransaction := request.ParseToTransaction(vaultId)
	if err = h.transactionService.PrepareTransaction(vault, &newTransaction); err != nil {
		h.logger.Errorf(err.Error())
		return dtos.TransactionDto{}, err
	}

	h.db.Transaction(func(tx *gorm.DB) error {
		vault.Total -= transaction.Value
		vault.Total += newTransaction.Value
		err = h.vaultRepository.Save(vault)
		if err != nil {
			h.logger.Errorf(err.Error())
			return err
		}

		err = h.transactionRepository.Update(&transaction, newTransaction)
		if err != nil {
			h.logger.Errorf(err.Error())
			return err
		}

		return nil
	})

	return transaction.ParseDto(), nil
}

func (h *transactionHandler) Get(vaultId, transactionId int) (dtos.TransactionDto, error) {
	_, transaction, err := h.getVaultAndTransaction(vaultId, transactionId)
	if err != nil {
		h.logger.Errorf(err.Error())
		return dtos.TransactionDto{}, err
	}

	return transaction.ParseDto(), nil
}

func (h *transactionHandler) GetAll(vaultId int, pagination utils.Pagination) ([]dtos.TransactionDto, error) {
	response := []dtos.TransactionDto{}

	transactions, err := h.transactionRepository.GetByVaultId(vaultId, pagination)
	if err != nil {
		return response, exceptions.VAULT_NOT_FOUND
	}

	for i := 0; i < len(transactions); i++ {
		transaction := transactions[i].ParseDto()
		response = append(response, transaction)
	}

	return response, nil
}

func (h *transactionHandler) getVaultAndTransaction(vaultId, transactionId int) (models.Vault, models.Transaction, error) {
	vault, err := h.vaultRepository.GetById(vaultId)
	if err != nil {
		return models.Vault{}, models.Transaction{}, exceptions.VAULT_NOT_FOUND
	}

	transaction, err := h.transactionRepository.GetById(transactionId)
	if err != nil {
		return models.Vault{}, models.Transaction{}, exceptions.TRANSACTION_NOT_FOUND
	}

	if transaction.Vault.ID != vault.ID {
		return models.Vault{}, models.Transaction{}, exceptions.TRANSACTION_NOT_BELONGS_TO_VAULT
	}

	return vault, transaction, err
}
