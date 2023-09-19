package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestTodo_FindOne(t *testing.T) {
	type args struct {
		ctx   context.Context
		input string
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
				ctx:   context.TODO(),
				input: "xxx",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(
						`SELECT 
    									id
									FROM todo`).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(nil).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
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

			result, err := ss.FindOne(test.args.ctx, test.args.input)
			fmt.Println(result, err)
			if !errors.Is(err, test.expectedError) {
				fmt.Println(err)
				t.Errorf("expected: %v but got: %v \n", test.expectedError, err)
			}
			fmt.Println("result", result)
		})
	}
}
