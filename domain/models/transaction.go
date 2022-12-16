package models

import "time"

type Transaction struct {
	ID          int
	VaultId     int `gorm:"column:vaultId"`
	Value       float32
	Date        time.Time
	Observation string
	CreatedAt   time.Time `gorm:"column:createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt"`

	// Relations
	Vault Vault
}
