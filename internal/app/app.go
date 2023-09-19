package app

import (
	"context"
	"github.com/sukha-id/bee/internal/app/config"
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

	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Error loading config file")
	}

	ctxLog := context.Background()
	logger := logrusx.NewProvider(&ctxLog, cfg.Log)

	db := initSqlConnection(&cfg)

	// init router
	r := initRouter(db, logger.GetLogger("bee-core"))

	// Create a server with desired configurations
	server := &http.Server{
		Addr:    "0.0.0.0:" + cfg.App.Port,
		Handler: r,
	}

	// Start the server in a separate goroutine
	go func() {
		logx.GetLogger().Info("Server running at: ", os.Getenv("APP_PORT"))
		if err := server.ListenAndServe(); err != nil {
			logger.GetLogger("bee-core").Fatal("", "Server error: %v", err)
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
