package handler_ping

import (
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/config"
	"github.com/sukha-id/bee/internal/app_rest/middleware/jwtx"
	"github.com/sukha-id/bee/internal/app_rest/service/service_auth"
	"github.com/sukha-id/bee/pkg/ginx"
	"github.com/sukha-id/bee/pkg/logrusx"
	"net/http"
)

type Handler struct {
	cfg         *config.ConfigApp
	authService service_auth.AuthServiceInterface
	jwtAuth     jwtx.AuthenticationInterface
	logger      *logrusx.LoggerEntry
}

func NewHandlerPing(
	router *gin.Engine,
	logger *logrusx.LoggerEntry) {

	v1 := router.Group("/api/v1/")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			var (
				guid = ctx.Value("request_id").(string)
			)
			logger.Info(guid, "ping")
			ginx.RespondWithJSON(ctx, http.StatusOK, "pong", nil)
		})
	}
}
