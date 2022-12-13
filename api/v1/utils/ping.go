package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlePing(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Pong")
}
