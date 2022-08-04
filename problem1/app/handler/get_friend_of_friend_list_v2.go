package handler

import (
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
	"net/http"
)

// NewGetFriendOfFriendListV2
//	友達の友達から友達とブロック関係のユーザーを除いたユーザー
//	user_id
//		基準のユーザーのid
//		正の数で出なければならない
func NewGetFriendOfFriendListV2(d dao.Dao) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := parseUserID(c.Param("user_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		us, err := d.User().GetFriendOfFriendListExceptBlockListAndFriendList(c.Request().Context(), userID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, us)
	}
}
