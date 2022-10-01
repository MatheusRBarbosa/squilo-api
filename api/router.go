package api

import (
	"github.com/gin-gonic/gin"
	v1Auth "github.com/matheusrbarbosa/gofin/api/v1/auth"
	v1User "github.com/matheusrbarbosa/gofin/api/v1/user"
	v1Utils "github.com/matheusrbarbosa/gofin/api/v1/utils"
)

func Router(server *gin.Engine) {
	v1Auth.RegisterAuthRoutes(server)
	v1Utils.RegisterUtilsRoutes(server)
	v1User.RegisterUserRoutes(server)
}
