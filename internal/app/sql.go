package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sukha-id/bee/pkg/logx"
	"os"
)

func initSqlConnection() *sqlx.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_DBNAME"))

	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		logx.GetLogger().Fatal(err)
	}

	return db
}
