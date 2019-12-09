package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() (left int) {
	for _, name := range users {
		coin := countCoin(name)
		distribution[name] = coin
		fmt.Printf("%s 的金币为：%d\n", name, coin)
	}
	result := 0
	for _, value := range distribution {
		result += value
	}
	fmt.Printf("总使用的金币： %d\n", result)
	return coins - result
}

func countCoin(name string) (result int) {
	if &name != nil {
		name = strings.ToUpper(name)
		bs := []byte(name)
		result = 0
		for _, value := range bs {
			if value == 'E' {
				result++
			} else if value == 'I' {
				result += 2
			} else if value == 'O' {
				result += 3
			} else if value == 'U' {
				result += 4
			}
		}
		return
	}
	return 0
}
