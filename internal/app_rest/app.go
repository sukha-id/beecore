package app_rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/config"
	"github.com/sukha-id/bee/database"
	"github.com/sukha-id/bee/internal/app_rest/handler/handler_auth"
	"github.com/sukha-id/bee/internal/app_rest/handler/handler_ping"
	"github.com/sukha-id/bee/internal/app_rest/middleware"
	"github.com/sukha-id/bee/internal/app_rest/middleware/jwtx"
	"github.com/sukha-id/bee/internal/app_rest/repositories/repo_auth"
	"github.com/sukha-id/bee/internal/app_rest/service/service_auth"
	"github.com/sukha-id/bee/pkg/logrusx"
	"github.com/sukha-id/bee/pkg/logx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Run(configPath string) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic("Error loading config file")
	}

	ctxLog := context.Background()
	logger := logrusx.NewProvider(&ctxLog, cfg.Log)

	mongoDB := database.InitMongoConnection(cfg)

	router := gin.Default()
	router.Use(middleware.RequestIDMiddleware(time.Duration(cfg.App.Timeout) * time.Second))

	handler_ping.NewHandlerPing(router, logger.GetLogger("monitoring"))

	repoAuth := repo_auth.NewAuthRepository(mongoDB, logger.GetLogger("repo-auth"))
	jwtAuthentication := jwtx.NewJWTAuthentication(cfg, repoAuth, logger.GetLogger("jwt-authorization"))
	serviceAuth := service_auth.NewAuthService(logger.GetLogger("service-auth"), repoAuth, jwtAuthentication)
	handler_auth.NewHandlerAuth(cfg, router, jwtAuthentication, serviceAuth, logger.GetLogger("handler-auth"))

	// Create a server with desired configurations
	server := &http.Server{
		Addr:    "0.0.0.0:" + cfg.App.Port,
		Handler: router,
	}

	// Now, set up the signal handling to catch SIGINT (Ctrl+C) and SIGTERM (kill)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		logx.GetLogger().Info("Server running at: ", cfg.App.Port)
		if err := server.ListenAndServe(); err != nil {
			logger.GetLogger("bee-core").Fatal("", "Server error", err)
		}
	}()

	// Block until a signal is received
	<-done

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		// close
		cancel()
	}()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		logger.GetLogger("bee-core").Fatal("", "Error during server shutdown: %v", err)
	}

	logger.GetLogger("bee-core").Info("", "Server gracefully shut down")
}
