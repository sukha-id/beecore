package repositories

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/sukha-id/bee/internal/domain/todo"
	"time"
)

func (t todo) StoreOne(ctx context.Context, todo domain.Todo) (id string, err error) {
	id = uuid.New().String()
	query := `INSERT INTO todo (id, task) VALUES (?,?,?)`
	_, err = t.db.ExecContext(ctx, query, id, todo.Task, time.Now())

	if err != nil {
		return
	}
	return
}
