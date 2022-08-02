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
		"users":  usersGenerator,
		"follow": followGenerator,
		"block":  blockGenerator,
	}
	n = 10000
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

func followGenerator() error {
	fmt.Print("INSERT INTO `follow_relation` (`following_user_id`, `followed_user_id`) VALUES")
	m := map[string]struct{}{}
	for i := 0; i < n; {
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

func blockGenerator() error {
	fmt.Print("INSERT INTO `block_relation` (`blocking_user_id`, `blocked_user_id`) VALUES")
	m := map[string]struct{}{}
	for i := 0; i < n; {
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
