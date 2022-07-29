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
	const query = `select 
    					users.id, 
    					users.user_id, 
    					users.name 
					from 
					    users 
					inner join 
					        (
					        	select 
					        		case 
					        	         when user1_id = ? then user2_id 
					        	         when user2_id = ? then user1_id 
					        	         else -1 end 
					        	    as 
					        	         user_id 
					        	 from 
					        	     friend_link 
					        	 where 
					        	     user1_id = ? 
					        	    or 
					        	     user2_id = ?
					         ) 
					            as 
					            friend_link_user_id 
					        on 
					            users.user_id = friend_link_user_id.user_id;`
	err := u.db.SelectContext(ctx, &users, query, userID, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUser(db *sqlx.DB) repository.User {
	return &User{db: db}
}
