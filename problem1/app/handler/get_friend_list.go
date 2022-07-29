package handler

import (
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
	"net/http"
)

func NewGetFriendList(d dao.Dao) echo.HandlerFunc {
	return func(c echo.Context) error {
		us, err := d.User().GetFriendList(c.Request().Context(), 1)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, us)
	}
}
