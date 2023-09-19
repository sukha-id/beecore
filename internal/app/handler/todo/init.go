package handler

import (
	domain "github.com/sukha-id/bee/internal/domain/todo"
)

type Handler struct {
	todoUseCase domain.TodoUseCase
}

func NewHandlerTodo(todoUseCase domain.TodoUseCase) Handler {
	return Handler{
		todoUseCase: todoUseCase,
	}
}
