package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var (
	generatorTypePtr = flag.String("type", "users", "choose generator type")
	generatorMap     = map[string]func() error{
		"users":  usersGenerator,
		"follow": followGenerator,
		"block":  blockGenerator,
	}
	nRelations = 50
	nUsers     = 10
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
	file, err := os.Open("/usr/share/dict/propernames")
	if err != nil {
		return err
	}
	defer file.Close()
	s := bufio.NewScanner(file)

	fmt.Print("INSERT INTO `users` (`user_id`, `name`) VALUES")
	for i := 0; s.Scan() && i < nUsers; i++ {
		query := fmt.Sprintf(`(%d, "%s")`, i+1, s.Text())
		if i == 0 {
			fmt.Print(" ", query)
		} else {
			fmt.Print(", ", query)
		}
	}
	fmt.Print(";")
	return nil
}

func followGenerator() error {
	fmt.Print("INSERT INTO `follow_relation` (`following_user_id`, `followed_user_id`) VALUES")
	m := map[string]struct{}{}
	for i := 0; i < nRelations; {
		a := rand.Int()%nUsers + 1
		b := rand.Int()%nUsers + 1
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
			fmt.Print(", ", query)
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
	for i := 0; i < nRelations; {
		a := rand.Int()%nUsers + 1
		b := rand.Int()%nUsers + 1
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
			fmt.Print(", ", query)
		}
		m[query] = struct{}{}
		i++
	}
	fmt.Print(";")
	return nil
}
