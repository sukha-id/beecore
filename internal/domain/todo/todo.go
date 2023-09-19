package domain

import (
	"context"
	"errors"
)

var (
	ERROR_TODO_HAS_EXIST = errors.New("todo has exist")
)

type TodoUseCase interface {
	StoreOne(ctx context.Context, todo Todo) (result Todo, err error)
}
