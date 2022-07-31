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

// GetFriendOfFriendListExceptBlockListAndFriendList go側でやるかsql側でやるか迷う　（ロジックのテストをデータベースないとできないのが微妙）
func (u *User) GetFriendOfFriendListExceptBlockListAndFriendList(ctx context.Context, userID int) ([]object.User, error) {
	const baseQuery = `
					with
					    block_relation as
					    (
					        select
					            CASE
					                WHEN blocking_user_id = :target_user_id THEN
					                    blocked_user_id
					                ELSE
					                    blocking_user_id
					                END as user_id
					        from
					            block_list
					        where
					            :target_user_id in (blocking_user_id, blocked_user_id)
					    ),
					    friend_list as
					    (
					        select
					            CASE
					                WHEN user1_id = :target_user_id THEN
					                    user2_id
					                ELSE
					                    user1_id
					                END as user_id
					        from
					            friend_link
					        where
					            :target_user_id in (user1_id, user2_id)
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
					    user_id in (select user_id from friend_of_friend_list)
					and
					    user_id not in (select user_id from friend_list)
					and
					    user_id not in (select user_id from block_relation);
`
	arg := map[string]interface{}{
		"target_user_id": userID,
	}
	query, args, err := sqlx.Named(baseQuery, arg)
	if err != nil {
		return nil, err
	}
	query = u.db.Rebind(query)
	var users []object.User
	err = u.db.SelectContext(ctx, &users, query, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) GetFriendOfFriendList(ctx context.Context, userID int) ([]object.User, error) {
	const baseQuery = `
					with
						friend_list as
						(
							select
								CASE
									WHEN user1_id = :target_user_id THEN
									    user2_id
									ELSE
									    user1_id
								END as user_id
							from
								friend_link
							where
								:target_user_id in (user1_id, user2_id)
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
	arg := map[string]interface{}{
		"target_user_id": userID,
	}
	query, args, err := sqlx.Named(baseQuery, arg)
	if err != nil {
		return nil, err
	}
	query = u.db.Rebind(query)
	var users []object.User
	err = u.db.SelectContext(ctx, &users, query, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) GetFriendList(ctx context.Context, userID int) ([]object.User, error) {
	const baseQuery = `
					with
					    friend_list as
						(
							select
								CASE
									WHEN user1_id = :target_user_id THEN
									    user2_id
									ELSE
									    user1_id
								END as user_id
							from
								friend_link
							where
								:target_user_id in (user1_id, user2_id)
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
	arg := map[string]interface{}{
		"target_user_id": userID,
	}
	query, args, err := sqlx.Named(baseQuery, arg)
	if err != nil {
		return nil, err
	}
	query = u.db.Rebind(query)
	var users []object.User
	err = u.db.SelectContext(ctx, &users, query, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUser(db *sqlx.DB) repository.User {
	return &User{db: db}
}
