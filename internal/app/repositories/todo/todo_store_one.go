package repositories

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/sukha-id/bee/internal/domain/todo"
	"time"
)

func (t *todo) StoreOne(ctx context.Context, todo domain.Task) (id string, err error) {
	var (
		guid = ctx.Value("request_id").(string)
	)
	id = uuid.New().String()
	query := `INSERT INTO todo (id, task) VALUES (?,?,?)`
	_, err = t.db.ExecContext(ctx, query, id, todo.Task, time.Now())

	if err != nil {
		t.logger.Error(guid, "error repository store one", err)
		return
	}
	return
}
