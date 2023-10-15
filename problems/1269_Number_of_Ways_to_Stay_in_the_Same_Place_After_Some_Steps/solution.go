package _269_Number_of_Ways_to_Stay_in_the_Same_Place_After_Some_Steps

// 解1 dp[k][i] = dp[k-1][i-1] + dp[k-1][i] + dp[k-1][i+1]
// 解2 -> 跑k步必須要返回，故i < K / 2 + 1 即可
func numWays(K int, n int) int {
	m := min(n, K/2+2)
	dp := make([]int, m)
	dp[0] = 1
	for k := 0; k < K; k++ {
		dp_t := make([]int, m)
		copy(dp_t, dp)
		for i := 0; i < m; i++ {
			dp[i] = dp_t[i]
			if i >= 1 {
				dp[i] += dp_t[i-1]
			}
			if i+1 < m {
				dp[i] += dp_t[i+1]
			}
			dp[i] %= int(1e9) + 7
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
