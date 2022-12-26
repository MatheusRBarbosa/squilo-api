package transaction

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	p "github.com/matheusrbarbosa/gofin/api/v1/utils/pagination"
	"github.com/matheusrbarbosa/gofin/application/handlers"
	"github.com/matheusrbarbosa/gofin/application/validators"
	l "github.com/matheusrbarbosa/gofin/crosscutting/logger"
	"github.com/matheusrbarbosa/gofin/domain/exceptions"
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
	vaultId, transactionId, err := getVaultAndTransactionIds(ctx)
	if err != nil {
		l.GetLogger().Error(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
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

	vaultId, transactionId, err := getVaultAndTransactionIds(ctx)
	if err != nil {
		l.GetLogger().Error(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
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

func handleGetTransaction(ctx *gin.Context) {
	vaultId, transactionId, err := getVaultAndTransactionIds(ctx)
	if err != nil {
		l.GetLogger().Error(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return

	}

	handler := handlers.TransactionHandler()
	response, err := handler.Get(vaultId, transactionId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

func handleGetAllTransactions(ctx *gin.Context) {
	vaultId, err := strconv.Atoi(ctx.Param("vaultId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
	}

	handler := handlers.TransactionHandler()
	response, err := handler.GetAll(vaultId, p.GetPagination(ctx))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

func getVaultAndTransactionIds(ctx *gin.Context) (int, int, error) {
	vaultId, err := strconv.Atoi(ctx.Param("vaultId"))
	if err != nil {
		return 0, 0, exceptions.URL_PARAMS_MISSING
	}

	transactionId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return 0, 0, exceptions.URL_PARAMS_MISSING
	}

	return vaultId, transactionId, nil
}
