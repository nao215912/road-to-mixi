package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
)

func UserIDPathParameter(c echo.Context) (int, error) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if userID < 0 {
		return 0, fmt.Errorf("id must be a positive number")
	}
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func LimitQueryParameter(c echo.Context) (int, error) {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return 0, err
	}
	if limit < 0 {
		return 0, fmt.Errorf("limit must be a positive number")
	}
	return limit, nil
}

func PageQueryParameter(c echo.Context) (int, error) {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, fmt.Errorf("page must be a positive number")
	}
	return page, nil
}
