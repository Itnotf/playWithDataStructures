package HJ4

import (
	"fmt"
)

func main() {
	var input string
	_, _ = fmt.Scan(&input)
	printEveryNChar(input)
}

func printEveryNChar(str string) {
	count := 0
	splitStr := [8]byte{'0', '0', '0', '0', '0', '0', '0', '0'}
	for _, val := range str {
		splitStr[count] = byte(val)
		if count == 7 {
			count = 0
			fmt.Println(string(splitStr[:]))
			splitStr = [8]byte{'0', '0', '0', '0', '0', '0', '0', '0'}
		} else {
			count++
		}
	}
	if count != 0 {
		lastStr := string(splitStr[:])
		fmt.Println(lastStr)
	}
}
