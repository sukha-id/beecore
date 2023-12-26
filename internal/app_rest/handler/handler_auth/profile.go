package handler_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/pkg/ginx"
	"net/http"
)

func (h *Handler) HandlerProfile(ctx *gin.Context) {
	var (
		guid     = ctx.Value("request_id").(string)
		userID   = ctx.GetString("user_id")
		username = ctx.GetString("username")
	)

	result, err := h.authService.Profile(ctx, userID, username)
	if err != nil {
		h.logger.Error(guid, "err", err)
		ginx.RespondWithJSON(ctx, http.StatusInternalServerError, "err", err.Error())
		return
	}

	ginx.RespondWithJSON(ctx, http.StatusOK, "success", result)
	return
}
