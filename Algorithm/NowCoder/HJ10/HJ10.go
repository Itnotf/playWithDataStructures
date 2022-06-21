package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	m := map[rune]bool{}

	for _, char := range str {
		m[char] = true
	}
	fmt.Println(len(m))
}
