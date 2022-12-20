package transaction

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/application/handlers"
	"github.com/matheusrbarbosa/gofin/application/validators"
	l "github.com/matheusrbarbosa/gofin/crosscutting/logger"
)

func handleCreateTransaction(context *gin.Context) {
	var request validators.CreateTransactionRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": validators.ParseError(err)})
		return
	}

	vaultId, err := strconv.Atoi(context.Param("vaultId"))
	if err != nil {
		l.GetLogger().Error(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid URL param"})
		return
	}

	handler := handlers.TransactionHandler()
	response, err := handler.Create(vaultId, request)
	if err != nil {
		context.Error(err)
		return
	}

	context.IndentedJSON(http.StatusCreated, response)
}
