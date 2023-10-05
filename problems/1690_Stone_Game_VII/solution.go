package _690_Stone_Game_VII

// 範例是錯的真的有夠搞
// stones = x x x x x x x
// dp[i][j] = stones[i:j]之間的最優策略 - 對手的最優策略
// 我方先手 所以dp[i][j] = max(s[i+1:j], s[i:j-1]) - dp[i+1][j]
// 設sum[i:j] = presum[j] - presum[i+1]
func stoneGameVII(stones []int) int {
	n := len(stones)
	stones = append([]int{0}, stones...)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	presum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		presum[i] = stones[i] + presum[i-1]
	}
	for i := 1; i <= n; i++ {
		dp[i][i] = 0
	}
	for i := 1; i+1 <= n; i++ {
		dp[i][i+1] = max(stones[i], stones[i+1])
	}
	for l := 3; l <= n; l++ {
		for i := 1; i+l-1 <= n; i++ {
			j := i + l - 1
			dp[i][j] = max(
				// 選nums[i]
				presum[j]-presum[i]-dp[i+1][j],
				presum[j-1]-presum[i-1]-dp[i][j-1],
			)
		}
	}
	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
