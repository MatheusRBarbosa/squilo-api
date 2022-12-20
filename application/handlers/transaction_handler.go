package handlers

import (
	"github.com/matheusrbarbosa/gofin/application/services"
	v "github.com/matheusrbarbosa/gofin/application/validators"
	"github.com/matheusrbarbosa/gofin/crosscutting/logger"
	"github.com/matheusrbarbosa/gofin/domain/dtos"
	"github.com/matheusrbarbosa/gofin/domain/exceptions"
	i "github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/infra/database"
	"github.com/matheusrbarbosa/gofin/infra/database/repositories"
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
