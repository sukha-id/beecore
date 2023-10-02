package todo

import "context"

type TodoRepository interface {
	FindOne(ctx context.Context, code string) (todo *Task, err error)
	StoreOne(ctx context.Context, todo Task) (uuid string, err error)
	UpdateOne(ctx context.Context, todo Task, uuid string) (err error)
}
