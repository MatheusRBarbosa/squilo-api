package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusrbarbosa/squilo/domain/exceptions"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, e := range c.Errors {
			err := e.Err
			if gErr, ok := err.(*exceptions.SquiloError); ok {
				c.JSON(gErr.Code, gin.H{
					"code":        gErr.Code,
					"message":     gErr.Message,
					"description": gErr.Description,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 500, "msg": "Server exception", "data": err.Error(),
				})
			}
		}
	}
}
