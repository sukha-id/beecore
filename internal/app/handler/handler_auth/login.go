package handler_auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sukha-id/bee/internal/app/service/service_auth"
	"github.com/sukha-id/bee/pkg/ginx"
	"github.com/sukha-id/bee/pkg/validatorx"
	"io"
	"net/http"
)

func (h *Handler) HandlerLogin(ctx *gin.Context) {
	var (
		guid  = ctx.Value("request_id").(string)
		param service_auth.LoginPayload
	)

	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&param); err != nil {
		h.logger.Error(guid, "err", err)
		ginx.RespondWithError(ctx, http.StatusInternalServerError, err.Error(), []string{err.Error()})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			h.logger.Error(guid, "err", err)
			ginx.RespondWithError(ctx, http.StatusInternalServerError, err.Error(), []string{err.Error()})
		}
	}(ctx.Request.Body)

	v := validator.New()
	if err := v.Struct(param); err != nil {
		h.logger.Error(guid, "err", err)
		ginx.RespondWithError(ctx, http.StatusUnprocessableEntity, err.Error(), validatorx.ExtractError(err))
		return
	}

	result, err := h.authService.Login(ctx, param)
	if err != nil {
		switch err.Error() {
		case "username or password is incorrect":
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
				http.StatusText(http.StatusUnauthorized),
				err.Error(),
			)
			return
		}
	}

	ginx.RespondWithJSON(ctx, http.StatusOK, "success", result)
	return
}
