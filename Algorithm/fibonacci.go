package algorithm

func fib(n int) int {
	const MOD int = 1e9 + 7

	dp := map[int]int{}
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % MOD
	}
	return dp[n]
}
