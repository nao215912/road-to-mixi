package dao

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"minimal_sns/configs"
	"minimal_sns/domain/object"
	"minimal_sns/domain/repository"
	"reflect"
	"sync"
	"testing"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func getDB() (*sqlx.DB, error) {
	var err error
	once.Do(func() {
		db, err = initDb(configs.Config{
			DB: configs.DBConfig{
				Driver:     "mysql",
				DataSource: "root:@(test_db:3306)/app",
			},
		})
	})
	return db, err
}

func newTestUser(insertQueries []string) (repository.User, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	err = deleteAll(db)
	if err != nil {
		return nil, err
	}
	err = insertAll(db, insertQueries)
	if err != nil {
		return nil, err
	}
	return NewUser(db), nil
}

func insertAll(db *sqlx.DB, insertQueries []string) error {
	for _, query := range insertQueries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteAll(db *sqlx.DB) error {
	for _, table := range []string{"users", "follow_relation", "block_relation"} {
		_, err := db.Exec(fmt.Sprintf("delete from %s", table))
		if err != nil {
			return err
		}
	}
	return nil
}

func TestGetFriendList(t *testing.T) {
	tests := []struct {
		name    string
		queries []string
		userID  int
		want    []object.User
		wantErr bool
	}{
		{
			name: "example",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 1), (1, 3), (3, 1)`,
			},
			userID: 1,
			want: []object.User{
				{
					ID:     2,
					UserID: 2,
					Name:   "Adam",
				},
				{
					ID:     3,
					UserID: 3,
					Name:   "Adlai",
				},
			},
			wantErr: false,
		},
		{
			name: "not_reciprocal_followings",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 1), (1, 3)`,
			},
			userID: 1,
			want: []object.User{
				{
					ID:     2,
					UserID: 2,
					Name:   "Adam",
				},
			},
			wantErr: false,
		},
		{
			name: "no_friend",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (1, 3)`,
			},
			userID:  1,
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := newTestUser(tt.queries)
			if err != nil {
				t.Fatal(err)
			}
			got, err := u.GetFriendList(context.Background(), tt.userID)
			if err != nil {
				t.Fatal(err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFriendOfFriendList(t *testing.T) {
	tests := []struct {
		name    string
		queries []string
		userID  int
		want    []object.User
		wantErr bool
	}{
		{
			name: "example",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 1), (2, 3), (3, 2)`,
			},
			userID: 1,
			want: []object.User{
				{
					ID:     1,
					UserID: 1,
					Name:   "Aaron",
				},
				{
					ID:     3,
					UserID: 3,
					Name:   "Adlai",
				},
			},
			wantErr: false,
		},
		{
			name: "no_friend_of_friend",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 1)`,
			},
			userID: 1,
			want: []object.User{
				{
					ID:     1,
					UserID: 1,
					Name:   "Aaron",
				},
			},
			wantErr: false,
		},
		{
			name: "no_friend",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 3), (3, 2)`,
			},
			userID:  1,
			want:    []object.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := newTestUser(tt.queries)
			if err != nil {
				t.Fatal(err)
			}
			got, err := u.GetFriendOfFriendList(context.Background(), tt.userID)
			if err != nil {
				t.Fatal(err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendOfFriendList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendOfFriendList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFriendOfFriendListExceptBlockListAndFriendList(t *testing.T) {
	tests := []struct {
		name    string
		queries []string
		userID  int
		want    []object.User
		wantErr bool
	}{
		{
			name: "example",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai"), (4, 4, "Adrian")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 1), (2, 3), (3, 2), (2, 4), (4, 2)`,
				`insert into block_relation (blocking_user_id, blocked_user_id) values (1, 4), (4, 1)`,
			},
			userID: 1,
			want: []object.User{
				{
					ID:     1,
					UserID: 1,
					Name:   "Aaron",
				},
				{
					ID:     3,
					UserID: 3,
					Name:   "Adlai",
				},
			},
			wantErr: false,
		},
		{
			name: "except_friend",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai"), (4, 4, "Adrian")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 1), (2, 3), (3, 2), (1, 3), (3, 1)`,
			},
			userID: 1,
			want: []object.User{
				{
					ID:     1,
					UserID: 1,
					Name:   "Aaron",
				},
			},
			wantErr: false,
		},
		{
			name: "except_block",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai"), (4, 4, "Adrian")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 1), (2, 3), (3, 2)`,
				`insert into block_relation (blocking_user_id, blocked_user_id) values (1, 3), (3, 1)`,
			},
			userID: 1,
			want: []object.User{
				{
					ID:     1,
					UserID: 1,
					Name:   "Aaron",
				},
			},
			wantErr: false,
		},
		{
			name: "no_friend",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai"), (4, 4, "Adrian")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 3), (3, 2)`,
			},
			userID:  1,
			want:    []object.User{},
			wantErr: false,
		},
		{
			name: "no_friend_of_friend",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai"), (4, 4, "Adrian")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 1), (3, 2)`,
			},
			userID: 1,
			want: []object.User{
				{
					ID:     1,
					UserID: 1,
					Name:   "Aaron",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := newTestUser(tt.queries)
			if err != nil {
				t.Fatal(err)
			}
			got, err := u.GetFriendOfFriendListExceptBlockListAndFriendList(context.Background(), tt.userID)
			if err != nil {
				t.Fatal(err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendOfFriendListExceptBlockListAndFriendList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendOfFriendListExceptBlockListAndFriendList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFriendListLimitOffset(t *testing.T) {
	tests := []struct {
		name    string
		queries []string
		userID  int
		limit   int
		offset  int
		want    []object.User
		wantErr bool
	}{
		{
			name: "example",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 1), (1, 3), (3, 1)`,
			},
			userID: 1,
			limit:  1,
			offset: 1,
			want: []object.User{
				{
					ID:     3,
					UserID: 3,
					Name:   "Adlai",
				},
			},
			wantErr: false,
		},
		{
			name: "limit",
			queries: []string{
				`insert into users (id, user_id, name) values (1, 1, "Aaron"), (2, 2, "Adam"), (3, 3, "Adlai")`,
				`insert into follow_relation (following_user_id, followed_user_id) values (1, 2), (2, 1), (1, 3), (3, 1)`,
			},
			userID: 1,
			limit:  1,
			offset: 0,
			want: []object.User{
				{
					ID:     2,
					UserID: 2,
					Name:   "Adam",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := newTestUser(tt.queries)
			if err != nil {
				t.Fatal(err)
			}
			got, err := u.GetFriendListLimitOffset(context.Background(), tt.userID, tt.limit, tt.offset)
			if err != nil {
				t.Fatal(err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendListLimitOffset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendListLimitOffset() got = %v, want %v", got, tt.want)
			}
		})
	}
}
