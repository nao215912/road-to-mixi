package dao

import (
	"github.com/jmoiron/sqlx"
	"problem1/configs"
	"problem1/domain/repository"
)

type Dao interface {
	User() repository.User
}

type dao struct {
	db *sqlx.DB
}

// Create DAO
func NewDao(conf configs.Config) (Dao, error) {
	db, err := initDb(conf)
	if err != nil {
		return nil, err
	}

	return &dao{db: db}, nil
}

func (d *dao) User() repository.User {
	return NewUser(d.db)
}
