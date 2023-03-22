package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(context *gin.Context, code int, message string) {
	context.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"error":   http.StatusText(code),
		"message": message,
	})
}

func ResponseSuccess(context *gin.Context, code int, data interface{}) {
	context.JSON(code, gin.H{
		"data": data,
	})
}
