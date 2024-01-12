package handler_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/internal/app_rest/service/service_auth"
	"github.com/sukha-id/bee/pkg/ginx"
	"github.com/sukha-id/bee/pkg/validatorx"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) HandlerSignUp(ctx *gin.Context) {
	var (
		guid  = ctx.Value("request_id").(string)
		param service_auth.SignUpPayload
	)

	cLogger := zap.L().With(
		zap.String("layer", "handler.signup"),
		zap.String("request_id", guid),
	)

	if err := ctx.ShouldBindJSON(&param); err != nil {
		cLogger.Error("error decode payload signup", zap.Error(err))
		ginx.RespondWithError(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil,
		)
		return
	}

	if err := validatorx.Validate(param); err != nil {
		cLogger.Warn("error validate payload signup", zap.Error(err))
		ginx.RespondWithError(ctx, http.StatusUnprocessableEntity, err.Error(), validatorx.ExtractError(err))
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
