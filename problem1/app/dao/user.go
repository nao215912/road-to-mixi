package dao

import (
	"github.com/jmoiron/sqlx"
	"problem1/domain/object"
	"problem1/domain/repository"
)

type User struct {
	db *sqlx.DB
}

func (u *User) GetFriendList(userID int) []object.User {
	//TODO implement me
	panic("implement me")
}

func NewUser(db *sqlx.DB) repository.User {
	return &User{db: db}
}
