package api

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/squilo/api/v1/auth"
	"github.com/matheusrbarbosa/squilo/api/v1/transaction"
	"github.com/matheusrbarbosa/squilo/api/v1/user"
	"github.com/matheusrbarbosa/squilo/api/v1/utils"
	"github.com/matheusrbarbosa/squilo/api/v1/vault"
)

func ApiRouter(server *gin.Engine) {
	router := server.Group("/api")
	v1 := router.Group("v1")

	auth.RegisterAuthRoutes(v1)
	utils.RegisterUtilsRoutes(v1)
	user.RegisterUserRoutes(v1)
	transaction.RegisterTransactionRoutes(v1)
	vault.RegisterVaultRoutes(v1)
}
