package ginx

import (
	"github.com/gin-gonic/gin"
)

type generalResponse struct {
	RequestID string      `json:"request_id,omitempty"`
	Message   string      `json:"message,omitempty"`
	Status    int         `json:"status,omitempty"`
	Error     interface{} `json:"error,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func RespondWithError(ctx *gin.Context, status int, message string, error interface{}) {
	ctx.JSON(status, generalResponse{
		RequestID: ctx.Value("request_id").(string),
		Message:   message,
		Status:    status,
		Error:     error,
	})
	ctx.Header("Content-Type", "application/json")
	ctx.Abort()
	return
}

func RespondWithJSON(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, generalResponse{
		RequestID: ctx.Value("request_id").(string),
		Message:   message,
		Status:    status,
		Data:      data,
	})
	ctx.Header("Content-Type", "application/json")
	ctx.Abort()
	return
}
