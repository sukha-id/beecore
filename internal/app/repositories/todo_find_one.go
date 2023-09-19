package repositories

import (
	"context"
	"database/sql"
	"errors"
	domainTodo "github.com/sukha-id/bee/internal/domain/todo"
)

func (t todo) FindOne(ctx context.Context, code string) (todo *domainTodo.Todo, err error) {
	var result domainTodo.Todo

	query := `SELECT
    			id
			FROM todo 
			WHERE id=?`

	err = t.db.SelectContext(ctx, &result, query, code)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}
