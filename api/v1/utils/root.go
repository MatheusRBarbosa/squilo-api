package utils

import (
	"github.com/gin-gonic/gin"
)

func RegisterUtilsRoutes(v1 *gin.RouterGroup) {
	v1.GET("ping", handlePing)
}
