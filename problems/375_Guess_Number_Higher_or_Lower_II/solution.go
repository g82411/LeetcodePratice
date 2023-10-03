package _75_Guess_Number_Higher_or_Lower_II

import "math"

// 區間
// dp[i][j] = min(dp[i][j], k + max(dp[i][k-1], dp[k+1][j])) for all k in [i, j]
func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for len := 2; len <= n; len++ {
		for i := 1; i+len-1 <= n; i++ {
			j := i + len - 1
			dp[i][j] = math.MaxInt / 2
			for k := i; k <= j; k++ {
				dp[i][j] = min(dp[i][j], k+max(dp[i][k-1], dp[k+1][j]))
			}
		}
	}
	return dp[1][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
