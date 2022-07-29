package dao

import (
	"context"
	"github.com/jmoiron/sqlx"
	"minimal_sns/domain/object"
	"minimal_sns/domain/repository"
)

type User struct {
	db *sqlx.DB
}

func (u *User) GetFriendList(ctx context.Context, userID int) ([]object.User, error) {
	var users []object.User
	err := u.db.SelectContext(ctx, &users, `select id, user_id, name from users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUser(db *sqlx.DB) repository.User {
	return &User{db: db}
}
