package _31_Minimum_Falling_Path_Sum

import "math"

func minFallingPathSum(m [][]int) int {
	width := len(m[0])
	height := len(m)
	dp := make([][]int, height)
	for i := range dp {
		dp[i] = make([]int, width)
		if i == 0 {
			for j := 0; j < width; j++ {
				dp[i][j] = m[i][j]
			}
		}
	}
	for i := 1; i < height; i++ {
		for j := 0; j < width; j++ {
			dp[i][j] = math.MaxInt
			for k := j - 1; k <= j+1; k++ {
				if k < 0 || k >= width {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+m[i][j])
			}
		}
	}
	res := math.MaxInt
	for i := 0; i < width; i++ {
		res = min(res, dp[height-1][i])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
