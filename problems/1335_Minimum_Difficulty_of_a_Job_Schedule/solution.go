package _335_Minimum_Difficulty_of_a_Job_Schedule

import "math"

// 關鍵點, 一個數組分成k分!!
// dp[i][j] => 在i天內完成j件任務的最大值
// TODO: 好像可以用monostack做
func minDifficulty(jd []int, d int) int {
	n := len(jd)
	if d > n {
		return -1
	}
	dp := make([][]int, n+1)
	for i := range dp {
		d := make([]int, d+1)
		for j := range d {
			d[j] = math.MaxInt / 2
		}
		dp[i] = d
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for k := 1; k <= d; k++ {
			maxD := 0
			for j := i - 1; j >= k-1; j-- {
				maxD = max(maxD, jd[j])
				dp[i][k] = min(dp[i][k], dp[j][k-1]+maxD)
			}
		}
	}
	return dp[n][d]
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
