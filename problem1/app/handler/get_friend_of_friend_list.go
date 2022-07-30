package handler

import (
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
	"net/http"
)

func NewGetFriendOfFriendList(d dao.Dao) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := UserIDPathParameter(c)
		if err != nil {
			return err
		}
		us, err := d.User().GetFriendOfFriendList(c.Request().Context(), userID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, us)
	}
}
