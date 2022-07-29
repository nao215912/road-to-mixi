package repository

import (
	"context"
	"minimal_sns/domain/object"
)

type User interface {
	GetFriendList(ctx context.Context, userID int) ([]object.User, error)
	GetFriendOfFriendList(ctx context.Context, userID int) ([]object.User, error)
}
