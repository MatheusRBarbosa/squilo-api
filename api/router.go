package api

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/api/v1/auth"
	"github.com/matheusrbarbosa/gofin/api/v1/transaction"
	"github.com/matheusrbarbosa/gofin/api/v1/user"
	"github.com/matheusrbarbosa/gofin/api/v1/utils"
)

func ApiRouter(server *gin.Engine) {
	router := server.Group("/api")
	v1 := router.Group("v1")

	auth.RegisterAuthRoutes(v1)
	utils.RegisterUtilsRoutes(v1)
	user.RegisterUserRoutes(v1)
	transaction.RegisterTransactionRoutes(v1)
}
