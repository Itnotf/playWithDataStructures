package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	m := map[int]int{}
	var keys []int
	for i := 0; i < n; i++ {
		var index int
		var value int
		fmt.Scan(&index)
		fmt.Scan(&value)

		if _, ok := m[index]; !ok {
			keys = append(keys, index)
		}
		m[index] += value
	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, m[v])
	}
}
