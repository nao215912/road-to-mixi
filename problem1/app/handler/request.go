package handler

import (
	"fmt"
	"strconv"
)

func parseUserID(s string) (int, error) {
	userID, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("user_id must be numeric")
	}
	if userID < 0 {
		return 0, fmt.Errorf("user_id must be a positive number")
	}
	return userID, nil
}

func parseLimit(s string) (int, error) {
	limit, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("limit must be numeric")
	}
	if limit < 0 {
		return 0, fmt.Errorf("limit must be a positive number")
	}
	return limit, nil
}

func parsePage(s string) (int, error) {
	page, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("page must be numeric")
	}
	if page < 0 {
		return 0, fmt.Errorf("page must be a positive number")
	}
	return page, nil
}
