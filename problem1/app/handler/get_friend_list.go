package handler

import (
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
	"net/http"
)

// NewGetFriendList
//	相互フォローのユーザー
//	user_id
//		基準のユーザーのid
//		正の数でなければいけない
func NewGetFriendList(d dao.Dao) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := parseUserID(c.Param("user_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		us, err := d.User().GetFriendList(c.Request().Context(), userID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, us)
	}
}
