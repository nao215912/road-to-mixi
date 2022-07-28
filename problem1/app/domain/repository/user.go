package repository

import "problem1/domain/object"

type User interface {
	GetFriendList(userID int) []object.User
}
