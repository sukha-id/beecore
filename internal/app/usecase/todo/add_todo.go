package usecase

import (
	"context"
	domain "github.com/sukha-id/bee/internal/domain/todo"
)

type todoUseCase struct {
}

func (t *todoUseCase) StoreOne(ctx context.Context, todo domain.Todo) (result domain.Todo, err error) {

	if todo.Task == "" {
		err = domain.ErrorTodoInvalidTask
		return
	}
	if false {
		err = domain.ErrorTodoHasExist
		return
	}

	return domain.Todo{}, nil
}

func NewTodoUseCase(repoTodo domain.TodoRepository) domain.TodoUseCase {
	return &todoUseCase{}
}
