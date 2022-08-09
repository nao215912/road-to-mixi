package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"minimal_sns/configs"
	"minimal_sns/domain/repository"
	"os"
	"testing"
)

var db *sqlx.DB

func newTestUser(insertQueries []string) (repository.User, error) {
	err := deleteAll(db)
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

func TestMain(m *testing.M) {

	err := os.Setenv("DB_DRIVER", "mysql")
	if err != nil {
		log.Fatalln(err)
	}
	err = os.Setenv("DB_DATASOURCE", "root:@(test_db:3306)/app")
	if err != nil {
		log.Fatalln(err)
	}
	db, err = initDb(configs.Get())
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}
