package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/config"
	"github.com/sukha-id/bee/internal/app/connector"
	handler "github.com/sukha-id/bee/internal/app/handler/todo"
	"github.com/sukha-id/bee/internal/app/middleware"
	repositories "github.com/sukha-id/bee/internal/app/repositories/todo"
	usecase "github.com/sukha-id/bee/internal/app/usecase/todo"
	"github.com/sukha-id/bee/pkg/logrusx"
	"github.com/sukha-id/bee/pkg/logx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Run() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic("Error loading config file")
	}

	ctxLog := context.Background()
	logger := logrusx.NewProvider(&ctxLog, cfg.Log)

	db, err := connector.InitSqlConnection(&cfg)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(middleware.TimeoutMiddleware(time.Duration(cfg.App.Timeout) * time.Second))

	repoTodo := repositories.NewRepositoryTodo(db, logger.GetLogger("bee-core"))
	useCaseTodo := usecase.NewTodoUseCase(logger.GetLogger("bee-core"), repoTodo)
	handler.NewHandlerTodo(router, logger.GetLogger("bee-core"), useCaseTodo)

	// Create a server with desired configurations
	server := &http.Server{
		Addr:    "0.0.0.0:" + cfg.App.Port,
		Handler: router,
	}

	// Start the server in a separate goroutine
	go func() {
		logx.GetLogger().Info("Server running at: ", cfg.App.Port)
		if err := server.ListenAndServe(); err != nil {
			logger.GetLogger("bee-core").Fatal("", "Server error", err)
		}
	}()

	// Now, set up the signal handling to catch SIGINT (Ctrl+C) and SIGTERM (kill)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-stop

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		logger.GetLogger("bee-core").Fatal("", "Error during server shutdown: %v", err)
	}

	logger.GetLogger("bee-core").Info("", "Server gracefully shut down")
}
