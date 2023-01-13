package api

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/squilo/api/configs/cors"
	"github.com/matheusrbarbosa/squilo/api/middleware"
	"github.com/matheusrbarbosa/squilo/infra/database"
)

func StartHttpServer() {
	server := gin.New()
	server.Use(middleware.ErrorHandler())
	server.Use(cors.Default())
	database.ConnectDatabase()
	ApiRouter(server)

	server.Run(":8000")
}
