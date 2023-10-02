package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/sukha-id/bee/config"
	"github.com/sukha-id/bee/pkg/logrusx"
	"testing"
)

func TestTodo_FindOne(t *testing.T) {
	type args struct {
		ctx   context.Context
		input string
	}

	cfg, err := config.LoadConfig("../../../../config.yaml")
	require.NoError(t, err)
	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, "request_id", uuid.New().String())
	logger := logrusx.NewProvider(&ctxWithValue, cfg.Log)

	testCase := []struct {
		name          string
		args          args
		expectedError error
		beforeTest    func(sqlmock.Sqlmock)
	}{
		{
			name: "test",
			args: args{
				ctx:   ctxWithValue,
				input: "xxx",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(
						`SELECT
										id
									FROM todo`).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(sql.ErrNoRows).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
			},
			expectedError: nil,
		},
		{
			name: "test no row",
			args: args{
				ctx:   ctxWithValue,
				input: "yyy",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(
						`SELECT 
    									id
									FROM todo`).
					WithArgs("yyy").
					WillReturnError(sql.ErrNoRows).
					WillReturnRows(sqlmock.NewRows([]string{"id"}))
			},
			expectedError: nil,
		},
	}

	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			ss := NewRepositoryTodo(db, logger.GetLogger("bee-core-todo-repository"))

			if test.beforeTest != nil {
				test.beforeTest(mockSQL)
			}

			_, err := ss.FindOne(test.args.ctx, test.args.input)
			if !errors.Is(err, test.expectedError) {
				t.Errorf("expected: %v but got: %v \n", test.expectedError, err)
			}
		})
	}
}
