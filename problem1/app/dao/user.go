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

func (u *User) GetFriendListLimitOffset(ctx context.Context, userID, limit, offset int) ([]object.User, error) {
	const baseQuery = `
						with
						    following as
						    (
						        select
						        	followed_user_id as user_id
						        from
						            follow_relation
						        where
						            following_user_id = :target_user_id
						    ),
						    followed as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						        	followed_user_id = :target_user_id
						    ),
						    follow as
						    (
						        select
						            user_id
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
						order by 
						    user_id
						limit 
							:target_limit
						offset 
							:target_offset
`
	query, args, err := sqlx.Named(baseQuery, map[string]interface{}{
		"target_user_id": userID,
		"target_limit":   limit,
		"target_offset":  offset,
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

// GetFriendOfFriendListExceptBlockListAndFriendList  ロジックをgo側でやるかsql側でやるか迷う　（ロジックのテストをデータベースないとできないのが微妙）
func (u *User) GetFriendOfFriendListExceptBlockListAndFriendList(ctx context.Context, userID int) ([]object.User, error) {
	const baseQuery = `
						with
						    following as
						    (
						        select
						        	followed_user_id as user_id
						        from
						            follow_relation
						        where
						            following_user_id = :target_user_id
						    ),
						    followed as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						        	followed_user_id = :target_user_id
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
						            followed_user_id as user_id
						        from
						            follow_relation
						        where
						            following_user_id in (select user_id from follow)
						    ),
						    follow_of_followed as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						            followed_user_id in (select user_id from follow)
						    ),
						    follow_of_follow as
						    (
						        select
						        	user_id
						        from
						            follow_of_following
						        where
						            user_id in (select user_id from follow_of_followed)
						    ),
						    blocking as
						    (
						        select
						            blocked_user_id as user_id
						        from
						            block_relation
						        where
						            blocking_user_id = :target_user_id
						    ),
						    blocked as
						    (
						        select
						            blocking_user_id as user_id
						        from
						            block_relation
						        where
						            blocked_user_id = :target_user_id
						    ),
						    block as
						    (
						        select
						        	user_id
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
						order by 
						    user_id
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
						        	followed_user_id as user_id
						        from
						            follow_relation
						        where
						            following_user_id = :target_user_id
						    ),
						    followed as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						        	followed_user_id = :target_user_id
						    ),
						    follow as
						    (
						        select
						            user_id
						        from
						            following
						        where
						            user_id in (select user_id from followed)
						    ),
						    follow_of_following as
						    (
						        select
						            followed_user_id as user_id
						        from
						            follow_relation
						        where
						            following_user_id in (select user_id from follow)
						    ),
						    follow_of_followed as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						            followed_user_id in (select user_id from follow)
						    ),
						    follow_of_follow as
						    (
						        select
						        	user_id
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
						order by 
						    user_id
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
						        	followed_user_id as user_id
						        from
						            follow_relation
						        where
						            following_user_id = :target_user_id
						    ),
						    followed as
						    (
						        select
						            following_user_id as user_id
						        from
						            follow_relation
						        where
						        	followed_user_id = :target_user_id
						    ),
						    follow as
						    (
						        select
						            user_id
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
						order by 
						    user_id
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
