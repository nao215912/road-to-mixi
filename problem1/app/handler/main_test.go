package handler

import (
	"log"
	"minimal_sns/configs"
	"minimal_sns/dao"
	"net/http"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {

	go func() {
		err := os.Setenv("SERVER_PORT", "8080")
		if err != nil {
			log.Fatalln(err)
		}
		err = os.Setenv("DB_DRIVER", "mysql")
		if err != nil {
			log.Fatalln(err)
		}
		err = os.Setenv("DB_DATASOURCE", "root:@(test_db:3306)/app")
		if err != nil {
			log.Fatalln(err)
		}
		conf := configs.Get()
		d, err := dao.NewDao(conf)
		if err != nil {
			log.Fatalln(err)
		}

		e := NewRouter(d)
		e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
	}()
	for {
		_, err := http.Get("http://localhost:8080")
		if err == nil {
			break
		}
	}
	os.Exit(m.Run())
}
