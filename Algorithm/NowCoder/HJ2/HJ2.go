package HJ2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	str := input.Text()
	input.Scan()
	char := input.Text()

	fmt.Println(getCountOfChar(str, char))
}

func getCountOfChar(str string, char string) int {
	str = strings.ToLower(str)
	char = strings.ToLower(char)
	count := 0

	for i := 0; i < len(str); i++ {
		if str[i] == char[0] {
			count++
		}
	}
	return count
}
