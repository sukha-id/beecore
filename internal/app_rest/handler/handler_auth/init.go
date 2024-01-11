package handler_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/config"
	"github.com/sukha-id/bee/internal/app_rest/middleware/jwtx"
	"github.com/sukha-id/bee/internal/app_rest/service/service_auth"
)

type Handler struct {
	cfg         *config.ConfigApp
	authService service_auth.AuthServiceInterface
	jwtAuth     jwtx.AuthenticationInterface
}

func NewHandlerAuth(
	cfg *config.ConfigApp,
	router *gin.Engine,
	jwtAuth jwtx.AuthenticationInterface,
	authService service_auth.AuthServiceInterface) {
	handler := &Handler{
		cfg:         cfg,
		authService: authService,
		jwtAuth:     jwtAuth,
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
