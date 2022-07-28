package dao

import (
	"github.com/jmoiron/sqlx"
	"minimal_sns/domain/object"
	"minimal_sns/domain/repository"
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
