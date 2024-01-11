package handler_auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sukha-id/bee/internal/app_rest/service/service_auth"
	"github.com/sukha-id/bee/pkg/ginx"
	"github.com/sukha-id/bee/pkg/validatorx"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func (h *Handler) HandlerLogin(ctx *gin.Context) {
	var (
		guid  = ctx.Value("request_id").(string)
		param service_auth.LoginPayload
	)

	cLogger := zap.L().With(
		zap.String("layer", "handler.login"),
		zap.String("request_id", guid),
	)

	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&param); err != nil {
		cLogger.Error("error decode", zap.Error(err))
		ginx.RespondWithError(ctx, http.StatusInternalServerError, err.Error(), []string{err.Error()})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			cLogger.Error("error read", zap.Error(err))
			ginx.RespondWithError(ctx, http.StatusInternalServerError, err.Error(), []string{err.Error()})
		}
	}(ctx.Request.Body)

	v := validator.New()
	if err := v.Struct(param); err != nil {
		cLogger.Error("error validate", zap.Error(err))
		ginx.RespondWithError(ctx, http.StatusUnprocessableEntity, err.Error(), validatorx.ExtractError(err))
		return
	}

	result, err := h.authService.Login(ctx, param)
	if err != nil {
		switch err.Error() {
		case "username or password is incorrect":
			cLogger.Warn("username or password is incorrect", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnauthorized,
				http.StatusText(http.StatusUnauthorized),
				err.Error(),
			)
			return
		default:
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusInternalServerError,
				http.StatusText(http.StatusUnauthorized),
				err.Error(),
			)
			return
		}
	}

	cLogger.Info("success handler login")
	ginx.RespondWithJSON(ctx, http.StatusOK, "success", result)
	return
}
