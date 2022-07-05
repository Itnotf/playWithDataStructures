package Leetcode

func runningSum(nums []int) []int {
	sum := make([]int, len(nums))

	for index, num := range nums {
		if index == 0 {
			sum[index] = num
		} else {
			sum[index] = num + sum[index-1]
		}
	}
	return sum
}
