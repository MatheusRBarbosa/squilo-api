package models

import "github.com/matheusrbarbosa/gofin/domain/dtos"

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (user *User) ParseDto() dtos.UserDto {
	return dtos.UserDto{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
