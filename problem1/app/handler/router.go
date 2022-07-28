package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"problem1/dao"
)

func NewRouter(d dao.Dao) *echo.Echo {
	e := echo.New()
	e.GET("/get_friend_list", NewGetFriendList(d))
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "members GET")
	})

	return e
}
