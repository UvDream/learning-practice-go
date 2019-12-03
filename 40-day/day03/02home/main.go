package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 判断几个汉字
func main() {
	s1 := "北京是全国one"
	var count int
	for _, c := range s1 {
		fmt.Println(c)
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)
	// how do you do
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	fmt.Println(m1)

}
