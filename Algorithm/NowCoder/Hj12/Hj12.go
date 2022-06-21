package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	strArr := []byte(str)

	l := len(strArr)
	for i := 0; i < l/2; i++ {
		strArr[i], strArr[l-i-1] = strArr[l-i-1], strArr[i]
	}
	fmt.Println(string(strArr))
}
