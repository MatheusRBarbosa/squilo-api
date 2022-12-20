package transaction

import (
	"github.com/gin-gonic/gin"
	m "github.com/matheusrbarbosa/gofin/api/middleware"
)

func RegisterTransactionRoutes(v1 *gin.RouterGroup) {
	v1.POST("vault/:id/transaction", m.ValidateJWT(), handleCreateTransaction)
}
