package handler

import (
	"github.com/gin-gonic/gin"
	domain "github.com/sukha-id/bee/internal/domain/todo"
	"github.com/sukha-id/bee/pkg/logrusx"
	"net/http"
)

type Handler struct {
	todoUseCase domain.TodoUseCase
	logger      *logrusx.LoggerEntry
}

func NewHandlerTodo(router *gin.Engine, todoUseCase domain.TodoUseCase, logger *logrusx.LoggerEntry) {
	handler := &Handler{
		todoUseCase: todoUseCase,
		logger:      logger,
	}
	v1 := router.Group("/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.GET("create", handler.HandlerCreateTodo)
	}
}
