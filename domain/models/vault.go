package models

import (
	"time"

	"github.com/matheusrbarbosa/squilo/domain/dtos"
	"gorm.io/gorm"
)

type Vault struct {
	ID          int
	Name        string
	Description string
	Configs     string //TODO: coluna eh json, verificar como manipular
	Total       float32
	UserId      int       `gorm:"column:userId"`
	TypeId      int       `gorm:"column:typeId"`
	CreatedAt   time.Time `gorm:"column:createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt"`
	DeletedAt   gorm.DeletedAt

	// Relations
	User         User
	Type         VaultType
	Transactions []Transaction
}

func (v *Vault) ParseDto() dtos.VaultDto {
	return dtos.VaultDto{
		Id:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Total:       v.Total,
		TypeId:      v.TypeId,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		Type: dtos.IdName{
			Id:   v.Type.ID,
			Name: v.Type.Name,
		},
	}
}

func (vault *Vault) BeforeCreate(tx *gorm.DB) error {
	vault.CreatedAt = time.Now()
	vault.UpdatedAt = time.Now()
	return nil
}

func (vault *Vault) BeforeUpdate(tx *gorm.DB) error {
	vault.UpdatedAt = time.Now()
	return nil
}
