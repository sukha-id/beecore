package repositories

import (
	"github.com/jmoiron/sqlx"
	domainTodo "github.com/sukha-id/bee/internal/domain/todo"
)

type todo struct {
	db *sqlx.DB
}

func NewRepositoryTodo(db *sqlx.DB) domainTodo.TodoRepository {
	return &todo{db: db}
}
