package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/application/handlers"
	"github.com/matheusrbarbosa/gofin/application/validators"
)

func handleSignup(context *gin.Context) {
	var request validators.SignupRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": validators.ParseError(err)})
		return
	}

	context.IndentedJSON(http.StatusOK, handlers.HandleSignup(request))
}
