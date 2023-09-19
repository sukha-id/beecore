package handler

import (
	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/sukha-id/bee/internal/app/middleware"
	"github.com/sukha-id/bee/internal/domain/mocks"
	domain "github.com/sukha-id/bee/internal/domain/todo"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateTodo(t *testing.T) {
	var mockParam domain.Todo
	err := faker.FakeData(&mockParam)
	assert.NoError(t, err)

	mockUseCase := new(mocks.TodoUseCase)
	mockUseCase.On("StoreOne", mock.Anything, mock.Anything).Return(mockParam, err)

	handlerTodo := NewHandlerTodo(mockUseCase)
	r := gin.Default()
	r.Use(middleware.TimeoutMiddleware(5 * time.Second))
	v1 := r.Group("/v1")
	{
		v1.GET("create", handlerTodo.HandlerCreateTodo)
	}

	req, err := http.NewRequest(http.MethodGet, "/v1/create", nil)
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreatePing(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.TimeoutMiddleware(5 * time.Second))
	v1 := r.Group("/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	req, err := http.NewRequest(http.MethodGet, "/v1/ping", nil)
	require.NoError(t, err)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := `{"message":"pong"}`
	assert.Equal(t, expectedResponse, rec.Body.String())
}
