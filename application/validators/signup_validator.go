package validators

import "github.com/matheusrbarbosa/gofin/domain/models"

type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=6"`
}

func (r *SignupRequest) ParseToUser() models.User {
	return models.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
