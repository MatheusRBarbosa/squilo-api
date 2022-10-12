package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/api/middleware"
)

func RegisterUtilsRoutes(v1 *gin.RouterGroup) {
	v1.GET("ping", handlePing)
	v1.GET("check-jwt", middleware.ValidateJWT(), handleCheckJwt)
}
