package usecase

import (
	"context"
	domain "github.com/sukha-id/bee/internal/domain/todo"
	"github.com/sukha-id/bee/pkg/logrusx"
)

type TodoUseCase struct {
	repoTodo domain.TodoRepository
	logger   *logrusx.LoggerEntry
}

func (t *TodoUseCase) StoreOne(ctx context.Context, todo domain.Task) (result domain.Task, err error) {
	var (
		guid = ctx.Value("request_id").(string)
	)
	if todo.Task == "" {
		err = domain.ErrorTodoInvalidTask
		t.logger.Error(guid, "error use case store one", err)
		return
	}
	if false {
		err = domain.ErrorTodoHasExist
		t.logger.Error(guid, "error use case store one", err)
		return
	}

	return domain.Task{}, nil
}

func NewTodoUseCase(logger *logrusx.LoggerEntry, repoTodo domain.TodoRepository) domain.TodoUseCase {
	return &TodoUseCase{
		repoTodo: repoTodo,
		logger:   logger,
	}
}
