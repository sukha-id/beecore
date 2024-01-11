package ginx

import (
	"github.com/gin-gonic/gin"
)

type generalResponse struct {
	Message string      `json:"message,omitempty"`
	Status  int         `json:"status,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondWithError(ctx *gin.Context, status int, message string, error interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, generalResponse{
		Status:  status,
		Message: message,
		Error:   error,
	})
	ctx.Abort()
	return
}

func RespondWithJSON(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, generalResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
	ctx.Abort()
	return
}
