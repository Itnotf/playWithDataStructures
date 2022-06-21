package HJ3

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	_, _ = fmt.Scan(&input)
	count, _ := strconv.Atoi(input)

	arr := [501]int{}
	for i := 0; i < count; i++ {
		_, _ = fmt.Scan(&input)

		number, _ := strconv.Atoi(input)
		arr[number] = 1
	}

	for key, val := range arr {
		if val > 0 {
			fmt.Println(key)
		}
	}
}
