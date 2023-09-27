package _76_Out_of_Boundary_Paths

// 能從初始點踢到界外 就能從界外給我踢回來！！
func findPaths(rows int, cols int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, maxMove+1)
	for i := range dp {
		d := make([][]int, rows)
		for j := range d {
			d[j] = make([]int, cols)
		}
		dp[i] = d
	}
	mod := int(1e9) + 7
	for step := 1; step <= maxMove; step++ {
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				// 從dp[step][y][x]開始踢踢踢踢踢踢踢踢
				dir := [][]int{
					{1, 0},
					{0, 1},
					{-1, 0},
					{0, -1},
				}
				for _, d := range dir {
					dx := x + d[0]
					dy := y + d[1]
					if dx < 0 || dy < 0 || dy >= rows || dx >= cols {
						dp[step][y][x] = (dp[step][y][x] + 1) % mod
						continue
					}
					dp[step][y][x] = (dp[step][y][x] + dp[step-1][dy][dx]) % mod
				}
			}
		}
	}
	return dp[maxMove][startRow][startColumn]
}
