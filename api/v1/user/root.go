package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("api/v1/signup", handleSignup)
}
