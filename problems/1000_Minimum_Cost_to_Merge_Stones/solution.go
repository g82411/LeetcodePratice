package _000_Minimum_Cost_to_Merge_Stones

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) > 0 {
		return -1
	}
	dp := make([][]int, n)
	for i := range dp {
		d := make([]int, n)
		for j := range d {
			d[j] = int(1e9)
		}
		dp[i] = d
	}
	for i := range dp {
		dp[i][i] = 0
	}
	presum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + stones[i-1]
	}
	// 區間型的template
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			for m := i; m < j; m += k - 1 {
				dp[i][j] = min(dp[i][j], dp[i][m]+dp[m+1][j])
			}
			if (l-1)%(k-1) == 0 {
				dp[i][j] += presum[j+1] - presum[i]
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
