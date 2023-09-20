package repositories

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	domainTodo "github.com/sukha-id/bee/internal/domain/todo"
	"testing"
)

func TestTodo_StoreOne(t *testing.T) {
	type args struct {
		ctx   context.Context
		input domainTodo.Todo
	}

	testCase := []struct {
		name          string
		args          args
		expectedError error
		beforeTest    func(sqlmock.Sqlmock)
	}{
		{
			name: "test",
			args: args{
				ctx: context.TODO(),
				input: domainTodo.Todo{
					Task: "task",
				},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(`INSERT INTO todo`).
					WithArgs(sqlmock.AnyArg(), "task", AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
	}

	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			ss := NewRepositoryTodo(db)

			if test.beforeTest != nil {
				test.beforeTest(mockSQL)
			}

			_, err := ss.StoreOne(test.args.ctx, test.args.input)
			if !errors.Is(err, test.expectedError) {
				t.Errorf("expected %v got %v \n", test.expectedError, err)
			}
		})
	}
}
