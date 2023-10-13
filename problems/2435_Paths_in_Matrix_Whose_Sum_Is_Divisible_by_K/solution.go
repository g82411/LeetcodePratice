package _435_Paths_in_Matrix_Whose_Sum_Is_Divisible_by_K

func numberOfPaths(grid [][]int, k int) int {
	const M = 1e9 + 7
	rows := len(grid)
	cols := len(grid[0])
	dp := make([][][]int, rows)
	for i := range dp {
		dp[i] = make([][]int, cols)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}
	sum := 0
	for i := 0; i < rows; i++ {
		sum = (sum + grid[i][0]) % M
		dp[i][0][sum%k] = 1
	}
	sum = 0
	for i := 0; i < cols; i++ {
		sum = (sum + grid[0][i]) % M
		dp[0][i][sum%k] = 1
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			for r := 0; r < k; r++ {
				t := ((r-grid[i][j])%k + k) % k
				dp[i][j][r] = (dp[i-1][j][t] + dp[i][j-1][t]) % M
			}
		}
	}
	return dp[rows-1][cols-1][0]
}
