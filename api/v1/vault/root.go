package vault

import (
	"github.com/gin-gonic/gin"
	m "github.com/matheusrbarbosa/squilo/api/middleware"
)

func RegisterVaultRoutes(v1 *gin.RouterGroup) {
	v1.Handlers = append(v1.Handlers, m.ValidateJWT())
	v1.GET("vault", handleGetUserVaults)
}
