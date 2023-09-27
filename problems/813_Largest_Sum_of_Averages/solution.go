package _13_Largest_Sum_of_Averages

import "math"

// 關鍵字: 給一個數組 拆成k分 求.....
// dp[i][j] => 從array[0:i]之間 分成k份的解

func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	nums = append([]int{0}, nums...)
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = float64(math.MinInt) / 2
	}
	for i := 1; i <= n; i++ {
		for p := 1; p <= min(k, i); p++ {
			sum := float64(0)
			for j := i; j >= p; j-- {
				sum += float64(nums[j])
				// dp[i][p] = max(nums[0:j-1]分成k-1份 + j:i的average)
				dp[i][p] = max(dp[i][p], dp[j-1][p-1]+sum/float64(i-j+1))
			}
		}
	}
	ret := float64(0)
	for p := 0; p <= k; p++ {
		ret = max(ret, dp[n][p])
	}
	return ret
}

func max(a, b float64) float64 {
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
