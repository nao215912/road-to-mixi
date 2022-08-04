package repository

import (
	"context"
	"minimal_sns/domain/object"
)

// User
//	handlerに１体１で対応するmethodを生やしているのが微妙
type User interface {
	GetFriendList(ctx context.Context, userID int) ([]object.User, error)
	GetFriendListLimitOffset(ctx context.Context, userID, limit, offset int) ([]object.User, error)
	GetFriendOfFriendList(ctx context.Context, userID int) ([]object.User, error)
	GetFriendOfFriendListExceptBlockListAndFriendList(ctx context.Context, userID int) ([]object.User, error)
}
