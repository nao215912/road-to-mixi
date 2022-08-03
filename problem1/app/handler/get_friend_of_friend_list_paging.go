package handler

import (
	"github.com/labstack/echo/v4"
	"minimal_sns/dao"
	"net/http"
)

func NewGetFriendListPaging(d dao.Dao) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := ConvertUserID(c.Param("user_id"))
		if err != nil {
			return err
		}
		limit, err := ConvertLimit(c.QueryParam("limit"))
		if err != nil {
			return err
		}
		page, err := ConvertPageQuery(c.QueryParam("page"))
		if err != nil {
			return err
		}
		us, err := d.User().GetFriendListLimitOffset(c.Request().Context(), userID, limit, page*limit)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, us)
	}
}
