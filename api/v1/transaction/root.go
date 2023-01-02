package transaction

import (
	"github.com/gin-gonic/gin"
	m "github.com/matheusrbarbosa/squilo/api/middleware"
)

func RegisterTransactionRoutes(v1 *gin.RouterGroup) {
	v1.Handlers = append(v1.Handlers, m.ValidateJWT())
	v1.POST("vault/:vaultId/transaction", handleCreateTransaction)
	v1.DELETE("vault/:vaultId/transaction/:id", m.VaultRelated(), handleDeleteTransaction)
	v1.PUT("vault/:vaultId/transaction/:id", m.VaultRelated(), handleUpdateTransaction)
	v1.GET("vault/:vaultId/transaction", m.VaultRelated(), handleGetAllTransactions)
	v1.GET("vault/:vaultId/transaction/:id", m.VaultRelated(), handleGetTransaction)
}
