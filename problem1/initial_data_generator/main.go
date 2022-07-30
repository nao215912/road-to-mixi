package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

var (
	generatorTypePtr = flag.String("type", "users", "choose generator type")
	generatorMap     = map[string]func(){
		"users":       usersGenerator,
		"friend_link": friendLinkGenerator,
		"block_list":  blockLinkGenerator,
	}
)

func init() {
	flag.Parse()
}

func main() {
	if generator, ok := generatorMap[*generatorTypePtr]; ok {
		generator()
	} else {
		log.Fatalln("unexpected generator type")
	}
}

func usersGenerator() {
	res, err := http.Get("https://ideas.fandom.com/wiki/List_of_Gods")
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	doc.Find(".new").Each(func(index int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		fmt.Printf("INSERT INTO `users` (`user_id`,`name`) VALUES (%d,\"%s\");\n", index+1, text)
	})
}

func friendLinkGenerator() {
	am := map[int]int{}
	bm := map[int]int{}
	for i := 0; i < 400; {
		a := rand.Int()%400 + 1
		b := rand.Int()%400 + 1
		if a >= b {
			continue
		}
		va, aok := am[a]
		vb, bok := bm[b]
		if aok && bok && va == vb {
			continue
		}
		fmt.Printf("INSERT INTO `friend_link` (`user1_id`,`user2_id`) VALUES (%d,%d);\n", a, b)
		am[a] = i
		bm[b] = i
		i++
	}
}

func blockLinkGenerator() {
	am := map[int]int{}
	bm := map[int]int{}
	for i := 0; i < 400; {
		a := rand.Int()%400 + 1
		b := rand.Int()%400 + 1
		if a == b {
			continue
		}
		va, aok := am[a]
		vb, bok := bm[b]
		if aok && bok && va == vb {
			continue
		}
		fmt.Printf("INSERT INTO `block_list` (`blocking_user_id`,`blocked_user_id`) VALUES (%d,%d);\n", a, b)
		am[a] = i
		bm[b] = i
		i++
	}
}
