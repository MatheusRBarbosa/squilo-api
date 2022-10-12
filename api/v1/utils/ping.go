package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/application/services"
)

func handlePing(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Pong")
}

func handleCheckJwt(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, fmt.Sprintf("User: %v", services.AuthService().GetAuthUser().Email))
}
