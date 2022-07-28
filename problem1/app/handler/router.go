package handler

import (
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
	"net/http"
)

func NewRouter(d dao.Dao) *echo.Echo {
	e := echo.New()
	e.GET("/get_friend_list", NewGetFriendList(d))
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "members GET")
	})

	return e
}
