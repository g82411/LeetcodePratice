package _41_Cherry_Pickup

import "math"

// 這題可以跟Cherry Pickup II 一起看
// 一開始想的是從左上到右下很基本, 但注意他要回來, 這時候可以變成兩條路徑
// 然後隱性constriant 路徑1下右 跟路徑2的上左總合一致

func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := make([][][]int, n+1)
	// fk go的初始化真的好麻煩
	for i := range dp {
		d := make([][]int, n+1)
		for j := range dp {
			dd := make([]int, n+1)
			for k := range dp {
				dd[k] = math.MinInt
			}
			d[j] = dd
		}
		dp[i] = d
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for x := 1; x <= n; x++ {
				y := i + j - x
				// 簡單的越界判斷 下面比較難
				if y < 1 || y > n {
					continue
				}
				if grid[i-1][j-1] == -1 {
					continue
				}
				if grid[x-1][y-1] == -1 {
					continue
				}
				// 如果在原點, 直接加grid[0][0], 因為他是出發點
				if i == 1 && j == 1 && x == 1 {
					dp[i][j][x] = grid[0][0]
					continue
				}
				// 下面判斷dp
				dp[i][j][x] = max(dp[i][j][x], dp[i-1][j][x])
				dp[i][j][x] = max(dp[i][j][x], dp[i][j-1][x])
				dp[i][j][x] = max(dp[i][j][x], dp[i-1][j][x-1])
				dp[i][j][x] = max(dp[i][j][x], dp[i][j-1][x-1])

				dp[i][j][x] += grid[i-1][j-1]
				if !(x == i && y == j) {
					dp[i][j][x] += grid[x-1][y-1]
				}
			}
		}
	}
	return max(dp[n][n][n], 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
