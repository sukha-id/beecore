package response

import (
	"github.com/gin-gonic/gin"
)

func RespondWithError(ctx *gin.Context, status int, message string, error interface{}) {
	ctx.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"error":   error,
	})
}

func RespondWithJSON(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
