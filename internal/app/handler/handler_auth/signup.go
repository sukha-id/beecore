package handler_auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sukha-id/bee/internal/app/service/service_auth"
	"github.com/sukha-id/bee/pkg/ginx"
	"github.com/sukha-id/bee/pkg/validatorx"
	"io"
	"net/http"
)

func (h *Handler) HandlerSignUp(ctx *gin.Context) {
	var (
		guid  = ctx.Value("request_id").(string)
		param service_auth.SignUpPayload
	)
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&param); err != nil {
		h.logger.Error(guid, "err", err)
		fmt.Println(err)
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
		fmt.Println(err)

		ginx.RespondWithError(ctx, http.StatusUnprocessableEntity, err.Error(), validatorx.ExtractError(err))
		return
	}

	result, err := h.authService.SignUp(ctx, param)
	if err != nil {
		switch err.Error() {
		case "this username already exists":
			h.logger.Error(guid, "err", err)
			ginx.RespondWithJSON(ctx,
				http.StatusUnprocessableEntity,
				http.StatusText(http.StatusUnprocessableEntity),
				err.Error(),
			)
			return
		default:
			h.logger.Error(guid, "err", err)
			ginx.RespondWithJSON(ctx, http.StatusInternalServerError, "err", err.Error())
			return
		}

	}

	ginx.RespondWithJSON(ctx, http.StatusOK, "succsesdds", result)
	return
}
