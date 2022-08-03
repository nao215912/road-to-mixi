package handler

import (
	"net/http"
	"testing"
)

func TestGetFriendOfFriendListV2(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want int
	}{
		{
			name: "example",
			url:  "http://localhost:1323/get_friend_of_friend_list_v2/12",
			want: http.StatusOK,
		},
		{
			name: "non_numeric_user_id",
			url:  "http://localhost:1323/get_friend_of_friend_list_v2/a",
			want: http.StatusBadRequest,
		},
		{
			name: "minus_user_id",
			url:  "http://localhost:1323/get_friend_of_friend_list_v2/-12",
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
