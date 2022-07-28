package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
	"net/http"
)

func NewGetFriendList(d dao.Dao) echo.HandlerFunc {
	fmt.Println("hello dao")
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "members   GET")
	}
}
