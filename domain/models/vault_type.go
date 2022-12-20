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
	// Vaults []Vault
}

func (vt *VaultType) BeforeCreate(tx *gorm.DB) error {
	vt.CreatedAt = time.Now()
	vt.UpdatedAt = time.Now()
	return nil
}

func (vt *VaultType) BeforeUpdate(tx *gorm.DB) error {
	vt.UpdatedAt = time.Now()
	return nil
}
