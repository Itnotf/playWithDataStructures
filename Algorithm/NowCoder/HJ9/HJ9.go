package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	m := 0
	x := 10

	v := [10]int{}
	for n > 0 {
		q := n % x
		if v[q] == 0 {
			v[q]++
			m += n % x
			m *= 10
			continue
		}
		n = n / 10
	}
	fmt.Println(m / 10)
}
