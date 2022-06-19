package algorithm

import "math"

func permute(nums []int) [][]int {
	return backTrack(nums)
}

func backTrack(nums []int) [][]int {
	var permutations [][]int
	for i, num := range nums {
		if num == math.MaxInt {
			continue
		}

		nums[i] = math.MaxInt
		backNums := backTrack(nums)
		if backNums == nil {
			p := []int{num}
			permutations = append(permutations, p)
			nums[i] = num
			continue
		}
		for _, b := range backNums {
			p := []int{num}
			p = append(p, b...)
			permutations = append(permutations, p)
		}
		nums[i] = num
	}
	return permutations
}
