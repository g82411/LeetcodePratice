package _956_Tallest_Billboard

func tallestBillboard(rods []int) int {
	// 他有一點像背包問題
	// dp[diff] => diff = k 時 l的最大值
	const OFFSET = 5000
	dp := make([]int, 2*OFFSET+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[OFFSET] = 0
	for _, rob := range rods {
		temp := make([]int, 2*OFFSET+1)
		copy(temp, dp)
		for diff := -OFFSET; diff < OFFSET; diff++ {
			j := OFFSET + diff
			if temp[j] == -1 {
				continue
			}
			if diff+rob < OFFSET {
				dp[j+rob] = max(dp[j+rob], temp[j]+rob)
			}
			if diff-rob >= -OFFSET {
				dp[j-rob] = max(dp[j-rob], temp[j])
			}
		}
	}
	return dp[OFFSET+0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
