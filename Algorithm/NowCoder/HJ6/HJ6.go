package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			num /= i
			fmt.Printf("%d ", i)
		}
	}
	if num != 1 {
		fmt.Println(num)
	}
}
