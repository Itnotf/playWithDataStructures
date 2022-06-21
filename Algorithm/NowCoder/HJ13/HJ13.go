package main

import (
	"fmt"
	"io"
)

func main() {
	var str string
	var strArr []string
	for {
		_, err := fmt.Scan(&str)
		if err == io.EOF {
			break
		}
		strArr = append(strArr, str)
	}

	for i := 0; i < len(strArr)/2; i++ {
		strArr[i], strArr[len(strArr)-i-1] = strArr[len(strArr)-i-1], strArr[i]
	}
	newStr := ""
	for _, val := range strArr {
		newStr += val
		newStr += " "
	}
	fmt.Println(newStr)
}
