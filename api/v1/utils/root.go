package utils

import "github.com/gin-gonic/gin"

func RegisterUtilsRoutes(server *gin.Engine) {
	server.GET("ping", handlePing)
}
