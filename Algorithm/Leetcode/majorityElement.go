package Leetcode

func majorityElement(nums []int) int {
	m := map[int]int{}

	max := 0
	for _, num := range nums {
		m[num]++
		if m[num] > m[max] {
			max = num
		}
	}
	return max
}
