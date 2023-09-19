package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/internal/app/handler/response"
	domain "github.com/sukha-id/bee/internal/domain/todo"
	"net/http"
)

func (h *Handler) HandlerCreateTodo(ctx *gin.Context) {
	var param domain.Todo
	//decoder := json.NewDecoder(ctx.Request.Body)
	//if err := decoder.Decode(&param); err != nil {
	//	response.RespondWithError(ctx, http.StatusInternalServerError, err.Error(), []string{err.Error()})
	//	return
	//}
	//defer ctx.Request.Body.Close()
	//
	//v := validator.New()
	//
	//err := v.RegisterValidation("phone", IsValidPhoneNumber)
	//if err != nil {
	//	response.RespondWithError(ctx, http.StatusInternalServerError, err.Error(), []string{err.Error()})
	//	return
	//}
	//
	//if err := v.Struct(param); err != nil {
	//	response.RespondWithError(ctx, http.StatusUnprocessableEntity, err.Error(), validatorx.ExtractError(err))
	//	return
	//}
	//
	result, err := h.todoUseCase.StoreOne(ctx, param)
	if err != nil {
		response.RespondWithJSON(ctx, http.StatusInternalServerError, "err", err.Error())
	}
	//fmt.Println(result)
	response.RespondWithJSON(ctx, http.StatusOK, "success", result)
}
