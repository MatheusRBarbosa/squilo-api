package models

import (
	"time"

	"github.com/matheusrbarbosa/gofin/domain/dtos"
)

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

func (t *Transaction) ParseDto() dtos.TransactionDto {
	return dtos.TransactionDto{
		Id:          t.ID,
		VaultId:     t.VaultId,
		Date:        t.Date,
		Observation: t.Observation,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}