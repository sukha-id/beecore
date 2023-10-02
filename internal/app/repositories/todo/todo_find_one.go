package repositories

import (
	"context"
	"database/sql"
	"errors"
	domainTodo "github.com/sukha-id/bee/internal/domain/todo"
)

func (t *todo) FindOne(ctx context.Context, code string) (todo *domainTodo.Todo, err error) {
	var (
		result domainTodo.Todo
		guid   = ctx.Value("request_id").(string)
	)

	query := `SELECT
    			id
			FROM todo 
			WHERE id=?`

	err = t.db.GetContext(ctx, &result, query, code)

	if errors.Is(err, sql.ErrNoRows) {
		t.logger.Error(guid, "error repository todo find one", err)
		return nil, nil
	}

	if err != nil {
		t.logger.Error(guid, "error repository todo find one", err)
		return nil, err
	}

	return &result, nil
}
