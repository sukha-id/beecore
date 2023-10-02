package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/sukha-id/bee/internal/domain/mocks"
	domain "github.com/sukha-id/bee/internal/domain/todo"
	"github.com/sukha-id/bee/pkg/logrusx"
	"testing"
)

func TestAddTodo(t *testing.T) {
	mockRepoTodo := new(mocks.TodoRepository)

	mockParam := domain.Task{
		Task: "Task Baru",
	}

	ctxLog := context.Background()
	logger := logrusx.NewProvider(&ctxLog, logrusx.Config{
		Dir:       "",
		FileName:  "",
		MaxSize:   0,
		LocalTime: false,
		Compress:  false,
	})

	t.Run("Test Add Todo", func(t *testing.T) {
		ctx := context.Background()
		ctxWithValue := context.WithValue(ctx, "request_id", uuid.New().String())

		mockRepoTodo.On("StoreOne", ctxWithValue, mockParam).Return("xxx", nil)

		u := NewTodoUseCase(logger.GetLogger("bee-core-use-case-todo"), mockRepoTodo)

		_, err := u.StoreOne(ctxWithValue, mockParam)

		assert.NoError(t, err)
	})

	t.Run("Test Add Todo With Invalid Name", func(t *testing.T) {
		ctx := context.Background()
		ctxWithValue := context.WithValue(ctx, "request_id", uuid.New().String())

		mockRepoTodo.On("StoreOne", ctxWithValue, mockParam).Return("xxx", nil)

		u := NewTodoUseCase(logger.GetLogger("bee-core-use-case-todo"), mockRepoTodo)

		_, err := u.StoreOne(ctxWithValue, domain.Task{Task: ""})

		assert.EqualError(t, err, domain.ErrorTodoInvalidTask.Error())
	})
}
