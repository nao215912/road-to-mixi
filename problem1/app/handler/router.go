package handler

import (
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
)

func NewRouter(d dao.Dao) *echo.Echo {
	e := echo.New()
	e.GET("/get_friend_list/:user_id", NewGetFriendList(d))
	e.GET("/get_friend_of_friend_list/:user_id", NewGetFriendOfFriendList(d))
	e.GET("/get_friend_of_friend_list_v2/:user_id", NewGetFriendOfFriendListV2(d))
	e.GET("/get_friend_of_friend_list_paging/:user_id", NewGetFriendListPaging(d))

	return e
}
