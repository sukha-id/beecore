package domain

import (
	"context"
	"errors"
)

var (
	ErrorTodoHasExist    = errors.New("todo has exist")
	ErrorTodoInvalidTask = errors.New("invalid task name")
)

type TodoUseCase interface {
	StoreOne(ctx context.Context, todo Todo) (result Todo, err error)
}
