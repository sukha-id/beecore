package connector

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sukha-id/bee/internal/app/configuration"
	"time"
)

func InitSqlConnection(cfg *configuration.ConfigApp) (db *sqlx.DB, err error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.HostName,
		cfg.Database.Port,
		cfg.Database.DatabaseName)

	db, err = sqlx.Open("mysql", connectionString)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	db.SetMaxIdleConns(cfg.Database.MaxIdleConnection)
	db.SetMaxOpenConns(cfg.Database.MaxOpenConnection)
	db.SetConnMaxLifetime(time.Duration(cfg.Database.MaxLifetimeConnection) * time.Second)

	return
}
