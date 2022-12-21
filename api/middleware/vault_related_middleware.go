package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/application/services"
	l "github.com/matheusrbarbosa/gofin/crosscutting/logger"
	"github.com/matheusrbarbosa/gofin/domain/exceptions"
)

func VaultRelated() gin.HandlerFunc {
	return func(context *gin.Context) {
		authServuce := services.AuthService()
		user := authServuce.GetAuthUser()

		vaultId, err := strconv.Atoi(context.Param("vaultId"))
		if err != nil {
			l.GetLogger().Error(err)
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid URL param"})
			return
		}

		isRelated := user.Vault(vaultId) != nil
		if !isRelated {
			context.AbortWithStatusJSON(exceptions.VAULT_NOT_USER_RELATED.Code, exceptions.VAULT_NOT_USER_RELATED)
			return
		}
	}
}
