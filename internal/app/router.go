package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	handler "github.com/sukha-id/bee/internal/app/handler/todo"
	"github.com/sukha-id/bee/internal/app/middleware"
	"github.com/sukha-id/bee/internal/app/repositories"
	usecase "github.com/sukha-id/bee/internal/app/usecase/todo"
	"net/http"
	"time"
)

func initRouter(db *sqlx.DB) http.Handler {
	repoTodo := repositories.NewRepositoryTodo(db)
	useCaseTodo := usecase.NewTodoUseCase(repoTodo)
	handlerTodo := handler.NewHandlerTodo(useCaseTodo)
	r := gin.Default()
	r.Use(middleware.TimeoutMiddleware(5 * time.Second))
	v1 := r.Group("/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			time.Sleep(6 * time.Second)
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.GET("create", handlerTodo.HandlerCreateTodo)
	}
	return r.Handler()
}
