package _850_Minimum_Moves_to_Spread_Stones_Over_Grid

import "math"

func minimumMoves(grid [][]int) int {
	// 注意看題目 9*9
	// 找到所有0點然後跟他爆了
	var bruteForce func(remainZero int) int
	bruteForce = func(remainZero int) int {
		if remainZero == 0 {
			return 0
		}
		ans := math.MaxInt
		for r1 := 0; r1 < 3; r1++ {
			for c1 := 0; c1 < 3; c1++ {
				if grid[r1][c1] == 0 {
					for r2 := 0; r2 < 3; r2++ {
						for c2 := 0; c2 < 3; c2++ {
							if grid[r2][c2] > 1 {
								grid[r1][c1]++
								grid[r2][c2]--
								d := abs(r1-r2) + abs(c1-c2)
								ans = min(ans, d+bruteForce(remainZero-1))
								grid[r2][c2]++
								grid[r1][c1]--
							}
						}
					}
				}
			}
		}
		return ans
	}
	r := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i][j] == 0 {
				r++
			}
		}
	}
	return bruteForce(r)

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
