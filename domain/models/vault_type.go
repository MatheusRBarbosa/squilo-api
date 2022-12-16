package models

import (
	"time"

	"gorm.io/gorm"
)

type VaultType struct {
	ID                        int
	Name                      string
	Description               string
	AllowPositiveTransactions bool      `gorm:"column:allowPositiveTransactions"`
	AllowNegativeTransactions bool      `gorm:"column:allowNegativeTransactions"`
	CreatedAt                 time.Time `gorm:"column:createdAt"`
	UpdatedAt                 time.Time `gorm:"column:updatedAt"`
	DeletedAt                 gorm.DeletedAt

	// Relations
	Vaults []Vault
}
