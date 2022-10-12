package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(v1 *gin.RouterGroup) {
	v1.POST("login", handleLogin)
}
