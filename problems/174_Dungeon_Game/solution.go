package _74_Dungeon_Game

func calculateMinimumHP(d [][]int) int {
	rows := len(d)
	cols := len(d[0])
	dp := make([][]int, rows)
	for i := range dp {
		r := make([]int, cols)
		for j := range r {
			r[j] = 1
		}
		dp[i] = r
	}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i == rows-1 && j == cols-1 {
				dp[i][j] = 1
				continue
			}
			if i == rows-1 {
				dp[i][j] = max(dp[i][j+1]-d[i][j+1], 1)
				continue
			}
			if j == cols-1 {
				dp[i][j] = max(dp[i+1][j]-d[i+1][j], 1)
				continue
			}
			dp[i][j] = min(
				dp[i][j+1]-d[i][j+1],
				dp[i+1][j]-d[i+1][j],
			)
			dp[i][j] = max(dp[i][j], 1)
		}
	}
	dp[0][0] -= d[0][0]
	return max(dp[0][0], 1)
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
