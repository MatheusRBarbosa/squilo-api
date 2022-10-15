package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/matheusrbarbosa/gofin/domain/dtos"
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
