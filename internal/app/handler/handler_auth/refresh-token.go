package handler_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/pkg/ginx"
	"net/http"
)

func (h *Handler) HandlerRefreshToken(ctx *gin.Context) {
	var (
		guid         = ctx.Value("request_id").(string)
		refreshToken = ctx.Request.Header.Get("Refresh-Token")
	)

	if refreshToken == "" {
		ginx.RespondWithJSON(ctx, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	result, err := h.authService.RefreshToken(ctx, refreshToken)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			h.logger.Error(guid, "username or password is incorrect", err)
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnauthorized,
				http.StatusText(http.StatusUnauthorized),
				err.Error(),
			)
			return
		default:
			h.logger.Error(guid, "err", err)
			ginx.RespondWithJSON(
				ctx,
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				err.Error(),
			)
			return
		}
	}

	ginx.RespondWithJSON(ctx, http.StatusOK, "success", result)
}
