package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Pong")
}
