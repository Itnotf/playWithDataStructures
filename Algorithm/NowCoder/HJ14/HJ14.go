package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var word string
	var words []string
	for i := 0; i < n; i++ {
		fmt.Scan(&word)
		words = append(words, word)
	}
	words = sortWords(words)
	for _, w := range words {
		fmt.Println(w)
	}
}

func sortWords(words []string) []string {
	for i := 0; i < len(words); i++ {
		for j := i; j > 0; j-- {
			k := 0
			for k < len(words[j]) && k < len(words[j-1]) {
				if words[j][k] < words[j-1][k] {
					words[j], words[j-1] = words[j-1], words[j]
					break
				} else if words[j][k] == words[j-1][k] {
					k++
					if k == len(words[j]) {
						words[j], words[j-1] = words[j-1], words[j]
					}
				} else {
					break
				}
			}
		}
	}
	return words
}
