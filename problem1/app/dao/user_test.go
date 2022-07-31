package dao

import (
	"context"
	"minimal_sns/configs"
	"minimal_sns/domain/object"
	"reflect"
	"testing"
)

func TestGetFriendOfFriendListExceptBlockListAndFriendList(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		want    []object.User
		wantErr bool
	}{
		{
			name:   "sample",
			userID: 51,
			want: []object.User{
				{
					ID:     51,
					UserID: 51,
					Name:   "Tyr",
				},
			},
			wantErr: false,
		},
		{
			name:    "no_friend",
			userID:  12,
			want:    nil,
			wantErr: false,
		},
	}
	d, err := NewDao(configs.Config{
		DB: configs.DBConfig{
			Driver:     "mysql",
			DataSource: "root:@(test_db:3306)/app",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.User().GetFriendOfFriendListExceptBlockListAndFriendList(context.Background(), tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendOfFriendListExceptBlockListAndFriendList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendOfFriendListExceptBlockListAndFriendList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFriendOfFriendList(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		want    []object.User
		wantErr bool
	}{
		{
			name:   "sample",
			userID: 51,
			want: []object.User{
				{
					ID:     51,
					UserID: 51,
					Name:   "Tyr",
				},
			},
			wantErr: false,
		},
		{
			name:    "no_friend",
			userID:  12,
			want:    nil,
			wantErr: false,
		},
	}
	d, err := NewDao(configs.Config{
		DB: configs.DBConfig{
			Driver:     "mysql",
			DataSource: "root:@(test_db:3306)/app",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.User().GetFriendOfFriendList(context.Background(), tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendOfFriendListExceptBlockListAndFriendList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendOfFriendListExceptBlockListAndFriendList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFriendList(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		want    []object.User
		wantErr bool
	}{
		{
			name:   "sample",
			userID: 51,
			want: []object.User{
				{
					ID:     14,
					UserID: 14,
					Name:   "Hebe",
				},
			},
			wantErr: false,
		},
		{
			name:    "no_friend",
			userID:  12,
			want:    nil,
			wantErr: false,
		},
	}
	d, err := NewDao(configs.Config{
		DB: configs.DBConfig{
			Driver:     "mysql",
			DataSource: "root:@(test_db:3306)/app",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.User().GetFriendList(context.Background(), tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendOfFriendListExceptBlockListAndFriendList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendOfFriendListExceptBlockListAndFriendList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
