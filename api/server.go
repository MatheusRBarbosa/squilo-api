package api

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/squilo/api/middleware"
	"github.com/matheusrbarbosa/squilo/infra/database"
)

func StartHttpServer() {
	server := gin.New()
	server.Use(middleware.ErrorHandler())
	database.ConnectDatabase()
	ApiRouter(server)

	server.Run(":8000")
}
