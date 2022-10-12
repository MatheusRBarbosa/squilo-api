package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/application/services"
	"github.com/matheusrbarbosa/gofin/domain/exceptions"
	"github.com/matheusrbarbosa/gofin/domain/models"
	"github.com/matheusrbarbosa/gofin/infra/database/repositories"
)

func ValidateJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		const SCHEMA = "Bearer"
		header := context.GetHeader("Authorization")
		jwt := strings.Trim(header[len(SCHEMA):], " ")

		token, _ := services.AuthService().Validate(jwt)
		if token.Valid {
			claims := token.Claims.(*models.UserCustomClaims)
			fmt.Printf("id: %v", claims.Id)
			user, err := repositories.UserRepository().GetById(claims.ID)

			if err != nil {
				context.AbortWithError(exceptions.UNAUTHORIZED.Code, exceptions.UNAUTHORIZED)
			}

			services.AuthService().SetAuthUser(user)
		} else {
			context.AbortWithError(exceptions.UNAUTHORIZED.Code, exceptions.UNAUTHORIZED)
		}
	}
}