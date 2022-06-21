package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	count := 0

	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	fmt.Println(count)
}
