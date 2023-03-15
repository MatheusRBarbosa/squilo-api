package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(v1 *gin.RouterGroup) {

	v1.POST("user", handleSignup)
}
