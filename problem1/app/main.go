package main

import (
	"fmt"
	"log"
	"minimal_sns/configs"
	"minimal_sns/dao"
	"minimal_sns/handler"
	"strconv"
)

func main() {

	conf := configs.Get()
	fmt.Println(conf.DB.DataSource)
	d, err := dao.NewDao(conf)
	if err != nil {
		log.Fatalln(err)
	}

	e := handler.NewRouter(d)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}
