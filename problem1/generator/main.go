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
	generatorMap     = map[string]func() error{
		"users":       usersGenerator,
		"friend_link": friendLinkGenerator,
		"block_list":  blockListGenerator,
	}
)

func init() {
	flag.Parse()
}

func main() {
	if generator, ok := generatorMap[*generatorTypePtr]; ok {
		if err := generator(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln("unexpected generator type")
	}
}

func usersGenerator() error {
	res, err := http.Get("https://ideas.fandom.com/wiki/List_of_Gods")
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}
	doc.Find(".new").Each(func(index int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		fmt.Printf("INSERT INTO `users` (`user_id`,`name`) VALUES (%d,\"%s\");\n", index+1, text)
	})
	return nil
}

func friendLinkGenerator() error {
	fmt.Print("INSERT INTO `friend_link` (`user1_id`, `user2_id`) VALUES")
	m := map[string]struct{}{}
	for i := 0; i < 100; {
		a := rand.Int()%400 + 1
		b := rand.Int()%400 + 1
		if a >= b {
			continue
		}
		query := fmt.Sprintf("(%d, %d)", a, b)
		if _, ok := m[query]; ok {
			continue
		}
		if i == 0 {
			fmt.Print(" ", query)
		} else {
			fmt.Print(" ,", query)
		}
		m[query] = struct{}{}
		i++
	}
	fmt.Print(";")
	return nil
}

func blockListGenerator() error {
	fmt.Print("INSERT INTO `block_list` (`blocking_user_id`, `blocked_user_id`) VALUES")
	m := map[string]struct{}{}
	for i := 0; i < 100; {
		a := rand.Int()%400 + 1
		b := rand.Int()%400 + 1
		if a == b {
			continue
		}
		query := fmt.Sprintf("(%d, %d)", a, b)
		if _, ok := m[query]; ok {
			continue
		}
		if i == 0 {
			fmt.Print(" ", query)
		} else {
			fmt.Print(" ,", query)
		}
		m[query] = struct{}{}
		i++
	}
	fmt.Print(";")
	return nil
}
