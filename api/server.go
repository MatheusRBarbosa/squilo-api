package api

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/api/middleware"
	"github.com/matheusrbarbosa/gofin/infra/database"
)

func StartHttpServer() {
	server := gin.New()
	server.Use(middleware.ErrorHandler())
	database.ConnectDatabase()
	Router(server)

	server.Run(":8000")
}
