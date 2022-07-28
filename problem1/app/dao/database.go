package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"problem1/configs"
)

type DBConfig interface {
	FormatDSN() string
}

func initDb(conf configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(conf.DB.Driver, conf.DB.DataSource)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open failed: %w", err)
	}

	return db, nil
}
