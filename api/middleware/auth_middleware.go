package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/application/services"
)

func ValidateJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		const SCHEMA = "Bearer"
		header := context.GetHeader("Authorization")
		jwt := header[len(SCHEMA):]

		token, err := services.JWTService().Validate(jwt)
		if token.Valid {
			claims := token.Claims
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			context.AbortWithError(http.StatusUnauthorized, gin.Error{})
		}
	}
}
