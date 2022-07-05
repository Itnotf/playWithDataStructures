package Leetcode

func pivotIndex(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	leftSum := 0
	for index, num := range nums {
		//注意需要整除
		if leftSum == (sum-num)/2 && (sum-num)%2 == 0 {
			return index
		}
		leftSum += num
	}
	return -1
}
