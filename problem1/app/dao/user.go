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

// GetFriendOfFriendListExceptBlockListAndFriendList ロジックをgo側でやるかsql側でやるか迷う　（ロジックのテストをデータベースないとできないのが微妙）
func (u *User) GetFriendOfFriendListExceptBlockListAndFriendList(ctx context.Context, userID int) ([]object.User, error) {
	const baseQuery = `
						with
						    following as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						            followed_user_id = :target_user_id
						    ),
						    followed as
						    (
						        select
						            followed_user_id as user_id
						        from
						            follow_relation
						        where
						            following_user_id = :target_user_id
						    ),
						    follow as
						    (
						        select
						            following.user_id as user_id
						        from
						            following
						        where
						            user_id in (select user_id from followed)
						    ),
						    follow_of_following as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						            followed_user_id in (select user_id from follow)
						    ),
						    follow_of_followed as
						    (
						        select
						            followed_user_id as user_id
						        from
						            follow_relation
						        where
						            following_user_id in (select user_id from follow)
						    ),
						    follow_of_follow as
						    (
						        select
						            follow_of_following.user_id as user_id
						        from
						            follow_of_following
						        where
						            user_id in (select user_id from follow_of_followed)
						    ),
						    blocking as
						    (
						        select
						            blocking_user_id as user_id
						        from
						            block_relation
						        where
						            blocked_user_id = :target_user_id
						    ),
						    blocked as
						    (
						        select
						            blocked_user_id as user_id
						        from
						            block_relation
						        where
						            blocking_user_id = :target_user_id
						    ),
						    block as
						    (
						        select
						            blocking.user_id as user_id
						        from
						            blocking
						        where
						            user_id in (select user_id from blocked)
						    )
						select
						    id,
						    user_id,
						    name
						from
						    users
						where
						    user_id in (select user_id from follow_of_follow)
						and
						    user_id not in (select user_id from follow)
						and
						    user_id not in (select user_id from block)
`
	query, args, err := sqlx.Named(baseQuery, map[string]interface{}{
		"target_user_id": userID,
	})
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
						    following as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						                followed_user_id = :target_user_id
						    ),
						    followed as
						    (
						        select
						            followed_user_id as user_id
						        from
						            follow_relation
						        where
						                following_user_id = :target_user_id
						    ),
						    follow as
						    (
						        select
						            following.user_id as user_id
						        from
						            following
						        where
						            user_id in (select user_id from followed)
						    ),
						    follow_of_following as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						            followed_user_id in (select user_id from follow)
						    ),
						    follow_of_followed as
						    (
						        select
						            followed_user_id as user_id
						        from
						            follow_relation
						        where
						            following_user_id in (select user_id from following)
						        and
						            following_user_id in (select user_id from followed)
						    ),
						    follow_of_follow as
						    (
						        select
						            follow_of_following.user_id as user_id
						        from
						            follow_of_following
						        where
						            user_id in (select user_id from follow_of_followed)
						    )
						select
						    id,
						    user_id,
						    name
						from
						    users
						where
						    user_id in (select user_id from follow_of_follow)
`
	query, args, err := sqlx.Named(baseQuery, map[string]interface{}{
		"target_user_id": userID,
	})
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
						    following as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						            followed_user_id = :target_user_id
						    ),
						    followed as
						    (
						        select
						            followed_user_id as user_id
						        from
						            follow_relation
						        where
						            following_user_id = :target_user_id
						    ),
						    follow as
						    (
						        select
						            following.user_id as user_id
						        from
						            following
						        where
						            user_id in (select user_id from followed)
						    )
						select
						    id,
						    user_id,
						    name
						from
						    users
						where
						    user_id in (select user_id from follow)
`
	query, args, err := sqlx.Named(baseQuery, map[string]interface{}{
		"target_user_id": userID,
	})
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
