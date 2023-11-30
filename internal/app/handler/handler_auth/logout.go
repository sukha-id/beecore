package handler_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/pkg/ginx"
	"net/http"
)

func (h *Handler) HandlerLogout(ctx *gin.Context) {
	var (
		guid  = ctx.Value("request_id").(string)
		token = ctx.GetString("token")
	)

	err := h.authService.Logout(ctx, token)
	if err != nil {
		h.logger.Error(guid, "err", err)
		ginx.RespondWithJSON(ctx, http.StatusInternalServerError, "err", err.Error())
		return
	}

	ginx.RespondWithJSON(ctx, http.StatusOK, "success", nil)
	return
}
