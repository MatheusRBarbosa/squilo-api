package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/squilo/application/services"
	"github.com/matheusrbarbosa/squilo/domain/exceptions"
	"github.com/matheusrbarbosa/squilo/domain/models"
	"github.com/matheusrbarbosa/squilo/infra/database/repositories"
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

		authServuce := services.AuthService()
		token, _ := authServuce.Validate(jwt)
		if token.Valid {
			claims := token.Claims.(*models.UserCustomClaims)
			user, err := repositories.UserRepository().GetById(claims.ID)

			if err != nil {
				context.AbortWithStatusJSON(exceptions.UNAUTHORIZED.Code, exceptions.UNAUTHORIZED)
			}

			authServuce.SetAuthUser(user)
		} else {
			context.AbortWithStatusJSON(exceptions.UNAUTHORIZED.Code, exceptions.UNAUTHORIZED)
		}
	}
}
