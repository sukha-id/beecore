package domain

import "context"

type TodoRepository interface {
	FindOne(ctx context.Context, code string) (todo *Todo, err error)
	StoreOne(ctx context.Context, todo Todo) (uuid string, err error)
	UpdateOne(ctx context.Context, todo Todo, uuid string) (err error)
}
