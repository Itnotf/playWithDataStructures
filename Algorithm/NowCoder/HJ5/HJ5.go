package HJ5

import (
	"fmt"
	"math"
)

func main() {
	var input string
	_, _ = fmt.Scan(&input)

	fmt.Println(conv16To10(input))
}

func conv16To10(str string) int {
	num := 0
	pos := 0
	for i := len(str) - 1; i > 1; i-- {
		char := str[i]
		if char > 64 {
			char = char - 64 + 9
		} else {
			char = char - 48
		}

		num += int(char) * int(math.Pow(float64(16), float64(pos)))
		pos++
	}
	return num
}
