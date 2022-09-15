package api

import (
	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	server := gin.New()

	Router(server)

	server.Run(":8000")
}
