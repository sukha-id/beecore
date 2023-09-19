package domain

import (
	"context"
	"errors"
)

var (
	ERROR_TODO_HAS_EXIST    = errors.New("todo has exist")
	ERROR_TODO_INVALID_TASK = errors.New("invalid task name")
)

type TodoUseCase interface {
	StoreOne(ctx context.Context, todo Todo) (result Todo, err error)
}
