package _688_Knight_Probability_in_Chessboard

// dp 也848行, 反正可dfs大概也都可dp
func knightProbability(n int, k int, row int, column int) float64 {
	directions := [][2]int{
		{1, 2},
		{2, 1},
		{-1, 2},
		{-2, 1},
		{-1, -2},
		{-2, -1},
		{1, -2},
		{2, -1},
	}
	var prev [26][26]float64
	var curr [26][26]float64
	prev[row][column] = 1
	for m := 1; m <= k; m++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				curr[i][j] = 0
				for _, d := range directions {
					prevX, prevY := i+d[0], j+d[1]
					if prevX >= n || prevY >= n {
						continue
					}
					if prevX < 0 || prevY < 0 {
						continue
					}
					curr[i][j] += prev[prevX][prevY] * 0.125
				}
			}
		}
		curr, prev = prev, curr
	}
	var ret float64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ret += prev[i][j]
		}
	}
	return ret
}

// 可以dfs
// func knightProbability(n int, k int, row int, column int) float64 {
//     directions := [][2]int {
//         {1,2},
//         {2,1},
//         {-1,2},
//         {-2,1},
//         {-1,-2},
//         {-2,-1},
//         {1,-2},
//         {2,-1},
//     }
//     var dfs func(x, y, step int, prevProb float64) float64
//     var cache [26][26][105]float64
//     dfs = func(x, y, step int, prevProb float64) float64 {
//         if step > k {
//             return float64(prevProb)
//         }
//         if x < 0 || y < 0 || x >= n || y >= n {
//             return float64(0)
//         }
//         c := cache[x][y][step]
//         if c > 0 {
//             return c
//         }

//         for _, d := range directions {
//             dx, dy := x + d[0], y + d[1]
//             cache[x][y][step] += dfs(dx, dy, step + 1, float64(prevProb / 8.0))
//         }
//         return cache[x][y][step]
//     }
//     return dfs(row, column, 0, float64(1))
// }
