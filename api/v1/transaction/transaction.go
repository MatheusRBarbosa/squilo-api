package transaction

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/application/handlers"
	"github.com/matheusrbarbosa/gofin/application/validators"
	l "github.com/matheusrbarbosa/gofin/crosscutting/logger"
)

func handleCreateTransaction(ctx *gin.Context) {
	var request validators.CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": validators.ParseError(err)})
		return
	}

	vaultId, err := strconv.Atoi(ctx.Param("vaultId"))
	if err != nil {
		l.GetLogger().Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid URL param"})
		return
	}

	handler := handlers.TransactionHandler()
	response, err := handler.Create(vaultId, request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, response)
}

func handleDeleteTransaction(ctx *gin.Context) {
	vaultId, err := strconv.Atoi(ctx.Param("vaultId"))
	if err != nil {
		l.GetLogger().Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid URL param"})
		return
	}

	transactionId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		l.GetLogger().Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid URL param"})
		return
	}

	handler := handlers.TransactionHandler()
	response, err := handler.Delete(vaultId, transactionId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

func handleUpdateTransaction(ctx *gin.Context) {
	var request validators.CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": validators.ParseError(err)})
		return
	}

	vaultId, err := strconv.Atoi(ctx.Param("vaultId"))
	if err != nil {
		l.GetLogger().Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid URL param"})
		return
	}

	transactionId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		l.GetLogger().Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid URL param"})
		return
	}

	handler := handlers.TransactionHandler()
	response, err := handler.Update(vaultId, transactionId, request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
