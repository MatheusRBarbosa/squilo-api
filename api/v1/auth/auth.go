package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/application/handlers"
	"github.com/matheusrbarbosa/gofin/application/validators"
)

func handleLogin(context *gin.Context) {
	var request validators.LoginRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": validators.ParseError(err)})
		return
	}

	response, err := handlers.HandleLogin(context, request)
	if err != nil {
		context.Error(err)
		return
	}

	context.IndentedJSON(http.StatusOK, response)
}
