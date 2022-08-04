package handler

import (
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
	"net/http"
)

// NewGetFriendListPaging
//	あるページの相互フォローのユーザー
//	user_id
//		基準のユーザーのid
//		正の数で出なければならない
//	limit
//		１ページのユーザー数
//		正の数で出なければならない
//	page
//		ページ数
//		正の数で出なければならない
func NewGetFriendListPaging(d dao.Dao) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := parseUserID(c.Param("user_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		limit, err := parseLimit(c.QueryParam("limit"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		page, err := parsePage(c.QueryParam("page"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		us, err := d.User().GetFriendListLimitOffset(c.Request().Context(), userID, limit, page*limit)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, us)
	}
}
