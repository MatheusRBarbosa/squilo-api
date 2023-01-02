package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/squilo/application/handlers"
	"github.com/matheusrbarbosa/squilo/application/validators"
)

func handleLogin(context *gin.Context) {
	var request validators.LoginRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": validators.ParseError(err)})
		return
	}

	handler := handlers.AuthHandler()
	response, err := handler.HandleLogin(request)
	if err != nil {
		context.Error(err)
		return
	}

	context.IndentedJSON(http.StatusOK, response)
}
