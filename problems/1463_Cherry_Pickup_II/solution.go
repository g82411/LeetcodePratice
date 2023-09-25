package _463_Cherry_Pickup_II

import "math"

func cherryPickup(grid [][]int) int {
	// 走到最後一行的最大收益 -> 走道前一行的最大收益 + 最後一行的收益
	// 兩個機器人走兩個維度
	// 然後看到走迷宮的dp -> [I][J]大部分都是座標
	rows := len(grid)
	cols := len(grid[0])
	dp := make([][]int, cols)
	for i := range dp {
		d := make([]int, cols)
		for j := range d {
			d[j] = math.MinInt
		}
		dp[i] = d
	}
	dp[0][cols-1] = grid[0][0] + grid[0][cols-1]
	for row := 1; row < rows; row++ {
		// 注意這邊 如果copy(prev, dp)
		// 隊最外層是deep copy
		// 但對下面一層仍然是Showllo copy
		prev := make([][]int, cols)
		for i := range prev {
			prev[i] = make([]int, cols)
			copy(prev[i], dp[i])
		}
		for col1 := 0; col1 < cols; col1++ {
			for col2 := 0; col2 < cols; col2++ {
				// 掃描row 的每一個cols 注意兩個機器人都會各走一次
				// 下面要確定從哪邊走下來收益最大
				for r1 := col1 - 1; r1 <= col1+1; r1++ {
					if r1 < 0 || r1 >= cols {
						continue
					}
					for r2 := col2 - 1; r2 <= col2+1; r2++ {
						if r2 < 0 || r2 >= cols {
							continue
						}
						if col1 != col2 {
							dp[col1][col2] = max(dp[col1][col2], prev[r1][r2]+grid[row][col1]+grid[row][col2])
							continue
						}
						dp[col1][col2] = max(dp[col1][col2], prev[r1][r2]+grid[row][col1])
					}
				}
			}
		}
	}
	res := 0
	for i := 0; i < cols; i++ {
		for j := 0; j < cols; j++ {
			res = max(res, dp[i][j])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
