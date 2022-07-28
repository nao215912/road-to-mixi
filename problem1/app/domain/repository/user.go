package repository

import "minimal_sns/domain/object"

type User interface {
	GetFriendList(userID int) []object.User
}
