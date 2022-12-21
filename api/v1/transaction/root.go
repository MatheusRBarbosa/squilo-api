package transaction

import (
	"github.com/gin-gonic/gin"
	m "github.com/matheusrbarbosa/gofin/api/middleware"
)

func RegisterTransactionRoutes(v1 *gin.RouterGroup) {
	v1.Handlers = append(v1.Handlers, m.ValidateJWT())
	v1.POST("vault/:vaultId/transaction", handleCreateTransaction)
	v1.DELETE("vault/:vaultId/transaction/:id", m.VaultRelated(), handleDeleteTransaction)
}
