package handler

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func UserIDPathParameter(c echo.Context) (int, error) {
	c.Param("user_id")
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	return userID, err
}
