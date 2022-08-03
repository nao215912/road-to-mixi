package handler

import (
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
	"net/http"
)

func NewGetFriendOfFriendListV2(d dao.Dao) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := ConvertUserID(c.Param("user_id"))
		if err != nil {
			return err
		}
		us, err := d.User().GetFriendOfFriendListExceptBlockListAndFriendList(c.Request().Context(), userID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, us)
	}
}
