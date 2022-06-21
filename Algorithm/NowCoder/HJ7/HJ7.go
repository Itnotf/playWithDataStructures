package main

import (
	"fmt"
)

func main() {
	var input float32
	fmt.Scan(&input)

	intPart := float32(int(input))
	if input-intPart >= 0.5 {
		fmt.Println(intPart + 1)
	} else {
		fmt.Println(intPart)
	}
}
