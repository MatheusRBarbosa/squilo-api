package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/gofin/domain/utils"
)

func GetPagination(ctx *gin.Context) utils.Pagination {
	pagination := utils.Pagination{}

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		pagination.Page = 1
	} else {
		pagination.Page = page
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		pagination.Limit = 15
	} else {
		pagination.Limit = limit
	}

	pagination.Offset = (pagination.Page - 1) * pagination.Limit
	return pagination
}
