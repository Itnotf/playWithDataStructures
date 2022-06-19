package playWithDataStructures

import "math"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		m := math.MaxInt32
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			if i == coin {
				m = 1
				continue
			}

			m = min(m, dp[i-coin]+1)
		}
		dp[i] = m
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
