package handler

import (
	"net/http"
	"testing"
)

func TestGetFriendListPaging(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want int
	}{
		{
			name: "example",
			url:  "http://localhost:8080/get_friend_of_friend_list_paging/42?limit=42&page=42",
			want: http.StatusOK,
		},
		{
			name: "non_numeric_user_id",
			url:  "http://localhost:8080/get_friend_of_friend_list_paging/a?limit=42&page=42",
			want: http.StatusBadRequest,
		},
		{
			name: "minus_user_id",
			url:  "http://localhost:8080/get_friend_of_friend_list_paging/-42?limit=42&page=42",
			want: http.StatusBadRequest,
		},
		{
			name: "non_limit",
			url:  "http://localhost:8080/get_friend_of_friend_list_paging/42?page=42",
			want: http.StatusBadRequest,
		},
		{
			name: "non_numeric_limit",
			url:  "http://localhost:8080/get_friend_of_friend_list_paging/42?limit=a&page=42",
			want: http.StatusBadRequest,
		},
		{
			name: "minus_limit",
			url:  "http://localhost:8080/get_friend_of_friend_list_paging/42?limit=-42&page=42",
			want: http.StatusBadRequest,
		},
		{
			name: "non_page",
			url:  "http://localhost:8080/get_friend_of_friend_list_paging/42?limit=42",
			want: http.StatusBadRequest,
		},
		{
			name: "non_numeric_page",
			url:  "http://localhost:8080/get_friend_of_friend_list_paging/42?limit=42&page=a",
			want: http.StatusBadRequest,
		},
		{
			name: "minus_page",
			url:  "http://localhost:8080/get_friend_of_friend_list_paging/42?limit=-42&page=-42",
			want: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := http.Get(tt.url)
			if err != nil {
				t.Fatal(err)
			}
			got := res.StatusCode
			if got != tt.want {
				t.Errorf("%s got = %v, want %v", tt.url, got, tt.want)
			}
		})
	}
}
