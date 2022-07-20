package Leetcode

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	//转成字符串数组
	var numStr []string
	for _, num := range nums {
		numStr = append(numStr, strconv.Itoa(num))
	}
	sort.Slice(numStr, func(i, j int) bool {
		return numStr[i]+numStr[j] > numStr[j]+numStr[i]
	})

	r := ""
	for _, str := range numStr {
		if r == "0" {
			continue
		}
		r += str
	}
	return r
}
