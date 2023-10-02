package repositories

import (
	"github.com/jmoiron/sqlx"
	domainTodo "github.com/sukha-id/bee/internal/domain/todo"
	"github.com/sukha-id/bee/pkg/logrusx"
)

type todo struct {
	db     *sqlx.DB
	logger *logrusx.LoggerEntry
}

func NewRepositoryTodo(db *sqlx.DB, logger *logrusx.LoggerEntry) domainTodo.TodoRepository {
	return &todo{
		db:     db,
		logger: logger,
	}
}
