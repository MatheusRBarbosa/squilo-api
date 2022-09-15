package models

import (
	"time"

	"github.com/matheusrbarbosa/gofin/domain/dtos"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

func (user *User) ParseDto() dtos.UserDto {
	return dtos.UserDto{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
