package api

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/api/v1/user"
	"github.com/matheusrbarbosa/gofin/api/v1/utils"
)

func Router(server *gin.Engine) {
	utils.RegisterUtilsRoutes(server)
	user.RegisterUserRoutes(server)
}
