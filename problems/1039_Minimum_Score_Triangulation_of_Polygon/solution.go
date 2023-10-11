package _039_Minimum_Score_Triangulation_of_Polygon

import "math"

func minScoreTriangulation(v []int) int {
	n := len(v)
	dp := make([][]int, n)
	for i := range dp {
		d := make([]int, n)
		for j := range d {
			d[j] = math.MaxInt / 3
		}
		dp[i] = d
	}
	for i := 0; i+1 < n; i++ {
		dp[i][i+1] = 0
	}
	for l := 3; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j],
					dp[i][k]+v[i]*v[k]*v[j]+dp[k][j],
				)
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
