package interfaces

import "github.com/matheusrbarbosa/gofin/domain/models"

type VaultRepository interface {
	Create(vault models.Vault) models.Vault
	GetById(id int) (models.Vault, error)
}
