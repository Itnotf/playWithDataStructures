package Leetcode

func lengthOfLIS(nums []int) int {
	//dp[i] means LIS end of nums[i]
	dp := make([]int, len(nums))

	for i := range nums {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxLength := 1
	for _, value := range dp {
		if value > maxLength {
			maxLength = value
		}
	}
	return maxLength
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
