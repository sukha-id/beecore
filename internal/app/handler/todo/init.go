package handler

import (
	domain "github.com/sukha-id/bee/internal/domain/todo"
	"github.com/sukha-id/bee/pkg/logrusx"
)

type Handler struct {
	todoUseCase domain.TodoUseCase
	logger      *logrusx.LoggerEntry
}

func NewHandlerTodo(todoUseCase domain.TodoUseCase, logger *logrusx.LoggerEntry) Handler {
	return Handler{
		todoUseCase: todoUseCase,
		logger:      logger,
	}
}
