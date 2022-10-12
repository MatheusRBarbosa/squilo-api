package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlePing(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Pong")
}

func handleCheckJwt(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "User logged in")
}
