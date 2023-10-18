package _218_Maximum_Value_of_K_Coins_From_Piles

// dp[i][j] 前i個piles, 取j個硬幣的最大值

func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	pilesPresum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum := 0
		pilesPresum[i] = append(pilesPresum[i], sum)
		for j := 0; j < len(piles[i]); j++ {
			sum += piles[i][j]
			pilesPresum[i] = append(pilesPresum[i], sum)
		}
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			for t := 0; t <= min(len(piles[i]), j); t++ {
				if i == 0 {
					dp[i][j] = max(dp[i][j], pilesPresum[i][t])
					continue
				}
				dp[i][j] = max(dp[i][j], dp[i-1][j-t]+pilesPresum[i][t])
			}
		}
	}
	return dp[n-1][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
