package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/matheusrbarbosa/gofin/domain/dtos"
	"gorm.io/gorm"
)

type UserCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	BirthDate time.Time `gorm:"column:birthDate"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt

	// Relations
	Vaults []Vault
}

func (user *User) Vault(id int) *Vault {
	for _, v := range user.Vaults {
		if v.ID == id {
			return &v
		}
	}

	return nil
}

func (user *User) ParseDto() dtos.UserDto {
	return dtos.UserDto{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) error {
	user.UpdatedAt = time.Now()
	return nil
}
