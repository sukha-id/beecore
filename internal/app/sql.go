package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sukha-id/bee/internal/app/config"
	"github.com/sukha-id/bee/pkg/logx"
	"time"
)

func initSqlConnection(cfg *config.ConfigApp) *sqlx.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.HostName,
		cfg.Database.Port,
		cfg.Database.DatabaseName)

	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		logx.GetLogger().Fatal(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(1600 * time.Second)

	return db
}
