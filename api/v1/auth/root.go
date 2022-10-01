package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(server *gin.Engine) {
	server.POST("api/v1/login", handleLogin)
}
