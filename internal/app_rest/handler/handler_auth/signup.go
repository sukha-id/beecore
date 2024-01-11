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

func (h *Handler) HandlerSignUp(ctx *gin.Context) {
	var (
		guid  = ctx.Value("request_id").(string)
		param service_auth.SignUpPayload
	)

	cLogger := zap.L().With(
		zap.String("layer", "handler.refresh_token"),
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

	result, err := h.authService.SignUp(ctx, param)
	if err != nil {
		switch err.Error() {
		case "this username already exists":
			cLogger.Error("error this username already exist", zap.Error(err))
			ginx.RespondWithJSON(ctx,
				http.StatusUnprocessableEntity,
				http.StatusText(http.StatusUnprocessableEntity),
				err.Error(),
			)
			return
		default:
			cLogger.Error("err", zap.Error(err))
			ginx.RespondWithJSON(ctx, http.StatusInternalServerError, "err", err.Error())
			return
		}

	}

	cLogger.Error("success signup", zap.Error(err))
	ginx.RespondWithJSON(ctx, http.StatusOK, "succsesdds", result)
	return
}
