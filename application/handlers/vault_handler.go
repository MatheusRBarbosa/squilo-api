package handlers

import (
	"github.com/matheusrbarbosa/gofin/application/services"
	"github.com/matheusrbarbosa/gofin/crosscutting/logger"
	"github.com/matheusrbarbosa/gofin/domain/dtos"
	"github.com/matheusrbarbosa/gofin/domain/exceptions"
	i "github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/domain/utils"
	"github.com/matheusrbarbosa/gofin/infra/database"
	"github.com/matheusrbarbosa/gofin/infra/database/repositories"
	"gorm.io/gorm"
)

type vaultHandler struct {
	logger          logger.Logger
	db              *gorm.DB
	authService     i.AuthService
	vaultRepository i.VaultRepository
}

func VaultHandler() *vaultHandler {
	return &vaultHandler{
		logger:          logger.GetLogger(),
		db:              database.Context(),
		authService:     services.AuthService(),
		vaultRepository: repositories.VaultRepository(),
	}
}

func (h *vaultHandler) GetAll(pagination utils.Pagination) ([]dtos.VaultDto, error) {
	user := h.authService.GetAuthUser()
	response := []dtos.VaultDto{}

	vaults, err := h.vaultRepository.GetByUserId(user.ID, pagination)
	if err != nil {
		return response, exceptions.VAULT_NOT_FOUND
	}

	for i := 0; i < len(vaults); i++ {
		vault := vaults[i]
		response = append(response, vault.ParseDto())
	}

	return response, nil
}
