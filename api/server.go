package api

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/infra/database"
)

func StartHttpServer() {
	server := gin.New()

	database.ConnectDatabase()

	Router(server)

	server.Run(":8000")
}
