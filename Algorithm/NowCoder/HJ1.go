package NowCoder

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		fmt.Println(getLengthOfLastWord(str))
	}
}

func getLengthOfLastWord(str string) int {
	length := 0

	i := len(str) - 1
	for i >= 0 {
		if str[i] == ' ' {
			return length
		} else {
			length++
		}
		i--
	}
	return length
}
