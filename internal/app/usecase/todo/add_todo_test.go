package usecase

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/sukha-id/bee/internal/domain/mocks"
	domain "github.com/sukha-id/bee/internal/domain/todo"
	"testing"
)

func TestAddTodo(t *testing.T) {
	mockRepoTodo := new(mocks.TodoRepository)

	mockParam := domain.Todo{
		Task: "Task Baru",
	}

	t.Run("Test Add Todo", func(t *testing.T) {
		mockRepoTodo.On("StoreOne", context.TODO(), mockParam).Return("xxx", nil)

		u := NewTodoUseCase(mockRepoTodo)

		result, err := u.StoreOne(context.TODO(), mockParam)
		fmt.Println(result, err)

		assert.NoError(t, err)
	})

	t.Run("Test Add Todo With Invalid Name", func(t *testing.T) {

		mockRepoTodo.On("StoreOne", context.TODO(), mockParam).Return("xxx", nil)

		u := NewTodoUseCase(mockRepoTodo)

		result, err := u.StoreOne(context.TODO(), domain.Todo{Task: ""})
		fmt.Println(result, err)

		assert.EqualError(t, err, domain.ERROR_TODO_INVALID_TASK.Error())
	})
}
