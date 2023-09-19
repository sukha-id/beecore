package app

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sukha-id/bee/pkg/logx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Run() {
	logx.InitLogger("info", "")
	err := godotenv.Load()
	if err != nil {
		logx.GetLogger().Fatal("Error loading .env file")
	}

	db := initSqlConnection()

	// init router
	r := initRouter(db)

	// Create a server with desired configurations
	server := &http.Server{
		Addr:    "0.0.0.0:" + os.Getenv("APP_PORT"),
		Handler: r,
	}

	// Start the server in a separate goroutine
	go func() {
		logx.GetLogger().Info("Server running at: ", os.Getenv("APP_PORT"))
		if err := server.ListenAndServe(); err != nil {
			logx.GetLogger().Fatalf("Server error: %v", err)
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
		logx.GetLogger().Fatalf("Error during server shutdown: %v", err)
	}

	logx.GetLogger().Info("Server gracefully shut down")
}
