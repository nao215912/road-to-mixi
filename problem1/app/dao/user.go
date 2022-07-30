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

func (u *User) GetFriendOfFriendList(ctx context.Context, userID int) ([]object.User, error) {
	var users []object.User
	const query = `
					with
						friend_list as
						(
							select
								CASE
									WHEN user1_id = ? THEN
									    user2_id
									ELSE 
									    user1_id
								END  as user_id
							from
								friend_link
							where
								? in (user1_id, user2_id)
						),
						friend_of_friend_list as
						(
							select
								CASE
									WHEN user1_id in (select user_id from friend_list) THEN
									    user2_id
									ELSE 
										user1_id
								END as user_id
							from
								friend_link
							where
								user1_id in (select user_id from friend_list)
							or
								user2_id in (select user_id from friend_list)
						)
		
					select
						distinct
						id,
						user_id,
						name
					from
						users
					where
						user_id in (select user_id from friend_of_friend_list);
		`
	err := u.db.SelectContext(ctx, &users, query, userID, userID)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) GetFriendList(ctx context.Context, userID int) ([]object.User, error) {
	var users []object.User
	const query = `
					with
					    friend_list as
						(
							select
								CASE
									WHEN user1_id = ? THEN
									    user2_id
									ELSE 
									    user1_id
								END as user_id
							from
								friend_link
							where
								? in (user1_id, user2_id)
						)

					select
					    id,
					    user_id,
					    name
					from
					    users
					where
					    user_id in (select user_id from friend_list);
`
	err := u.db.SelectContext(ctx, &users, query, userID, userID)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUser(db *sqlx.DB) repository.User {
	return &User{db: db}
}
