package handler

import (
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
)

func NewRouter(d dao.Dao) *echo.Echo {
	e := echo.New()
	e.GET("/get_friend_list/:user_id", NewGetFriendList(d))

	return e
}
