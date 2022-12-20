package transaction

import (
	"github.com/gin-gonic/gin"
	m "github.com/matheusrbarbosa/gofin/api/middleware"
)

func RegisterTransactionRoutes(v1 *gin.RouterGroup) {
	v1.POST("vault/:vaultId/transaction", m.ValidateJWT(), handleCreateTransaction)
	v1.DELETE("vault/:vaultId/transaction/:id", m.ValidateJWT(), handleDeleteTransaction)
}
