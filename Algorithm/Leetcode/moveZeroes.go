package Leetcode

func moveZeroes(nums []int) {
	//双指针
	//i 指遍历到的当前位置
	//j 指0开始的位置
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j = i
			break
		}

		if i == len(nums)-1 {
			return
		}
	}

	for i := j + 1; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[j] = nums[i]
		nums[i] = 0
		j++
	}
}
