package services

import (
	"time"

	vtypes "github.com/matheusrbarbosa/gofin/domain/enums/vault_types"
	"github.com/matheusrbarbosa/gofin/domain/interfaces"
	"github.com/matheusrbarbosa/gofin/domain/models"
)

type userService struct{}

func UserService() interfaces.UserService {
	return &userService{}
}

func (s *userService) PrepareToCreate(user *models.User) {
	defaultVault := models.Vault{
		Name:        "Geral",
		Description: "Cofrinho inicial",
		Configs:     "{}",
		TypeId:      vtypes.General,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	user.Vaults = append(user.Vaults, defaultVault)
}
