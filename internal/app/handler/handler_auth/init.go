package handler_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/config"
	"github.com/sukha-id/bee/internal/app/middleware/jwtx"
	"github.com/sukha-id/bee/internal/app/service/service_auth"
	"github.com/sukha-id/bee/pkg/logrusx"
)

type Handler struct {
	cfg         *config.ConfigApp
	authService service_auth.AuthServiceInterface
	jwtAuth     jwtx.AuthenticationInterface
	logger      *logrusx.LoggerEntry
}

func NewHandlerAuth(
	cfg *config.ConfigApp,
	router *gin.Engine,
	jwtAuth jwtx.AuthenticationInterface,
	authService service_auth.AuthServiceInterface,
	logger *logrusx.LoggerEntry) {
	handler := &Handler{
		cfg:         cfg,
		authService: authService,
		jwtAuth:     jwtAuth,
		logger:      logger,
	}
	v1 := router.Group("/api/v1/auth")
	{
		v1.POST("/login", handler.HandlerLogin)
		v1.POST("/signup", handler.HandlerSignUp)
		v1.GET("/refresh-token", handler.HandlerRefreshToken)

	}
	v1WithAuth := router.Group("/api/v1/auth")
	{
		v1WithAuth.Use(jwtAuth.Authentication())
		v1WithAuth.GET("/profile", handler.HandlerProfile)
		v1WithAuth.GET("/logout", handler.HandlerLogout)
	}
}
