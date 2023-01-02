package models

import (
	"time"

	"github.com/matheusrbarbosa/squilo/domain/dtos"
	"gorm.io/gorm"
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
		Value:       t.Value,
		Date:        t.Date,
		Observation: t.Observation,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func (tr *Transaction) BeforeCreate(tx *gorm.DB) error {
	tr.CreatedAt = time.Now()
	tr.UpdatedAt = time.Now()
	return nil
}

func (tr *Transaction) BeforeUpdate(tx *gorm.DB) error {
	tr.UpdatedAt = time.Now()
	return nil
}
