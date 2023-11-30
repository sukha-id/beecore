package handler

import (
	"github.com/gin-gonic/gin"
	domain "github.com/sukha-id/bee/internal/domain/todo"
	"github.com/sukha-id/bee/pkg/logrusx"
)

type Handler struct {
	todoUseCase domain.TodoUseCase
	logger      *logrusx.LoggerEntry
}

func NewHandlerTodo(router *gin.Engine, logger *logrusx.LoggerEntry, todoUseCase domain.TodoUseCase) {
	handler := &Handler{
		todoUseCase: todoUseCase,
		logger:      logger,
	}
	v1 := router.Group("/api/v1/todo")
	{
		v1.GET("create", handler.HandlerCreateTodo)
	}
}
