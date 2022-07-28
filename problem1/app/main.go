package main

import (
	"fmt"
	"log"
	"problem1/configs"
	"problem1/dao"
	"problem1/handler"
	"strconv"
)

func main() {

	conf := configs.Get()
	d, err := dao.NewDao(conf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("hoge hoge")

	e := handler.NewRouter(d)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}
