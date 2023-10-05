package _246_Palindrome_Removal

import "math"

func minimumMoves(arr []int) int {
	n := len(arr)

	arr = append([]int{0}, arr...)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for l := 1; l <= n; l++ {
		for i := 1; i+l-1 <= n; i++ {
			j := i + l - 1
			dp[i][j] = math.MaxInt / 2
			for k := i; k <= j; k++ {
				if arr[k] == arr[j] {
					dp[i][j] = min(dp[i][j], dp[i][k-1]+max(dp[k+1][j-1], 1))
				}
			}
		}
	}
	return dp[1][n]
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
