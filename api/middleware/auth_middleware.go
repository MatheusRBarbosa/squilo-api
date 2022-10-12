package middleware

import (
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
		if header == "" {
			context.AbortWithStatusJSON(exceptions.UNAUTHORIZED.Code, exceptions.UNAUTHORIZED)
			return
		}

		jwt := strings.Trim(header[len(SCHEMA):], " ")

		token, _ := services.AuthService().Validate(jwt)
		if token.Valid {
			claims := token.Claims.(*models.UserCustomClaims)
			user, err := repositories.UserRepository().GetById(claims.ID)

			if err != nil {
				context.AbortWithStatusJSON(exceptions.UNAUTHORIZED.Code, exceptions.UNAUTHORIZED)
			}

			services.AuthService().SetAuthUser(user)
		} else {
			context.AbortWithStatusJSON(exceptions.UNAUTHORIZED.Code, exceptions.UNAUTHORIZED)
		}
	}
}
