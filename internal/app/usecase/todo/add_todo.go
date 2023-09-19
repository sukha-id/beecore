package usecase

import (
	"context"
	domain "github.com/sukha-id/bee/internal/domain/todo"
)

type todoUseCase struct {
}

func (t *todoUseCase) StoreOne(ctx context.Context, todo domain.Todo) (result domain.Todo, err error) {

	if todo.Task == "" {
		err = domain.ERROR_TODO_INVALID_TASK
		return
	}
	if false {
		err = domain.ERROR_TODO_HAS_EXIST
		return
	}

	return domain.Todo{}, nil
}

func NewTodoUseCase(repoTodo domain.TodoRepository) domain.TodoUseCase {
	return &todoUseCase{}
}
