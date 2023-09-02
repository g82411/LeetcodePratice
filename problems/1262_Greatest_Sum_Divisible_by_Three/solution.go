package _262_Greatest_Sum_Divisible_by_Three

import (
	"math"
)

func maxSumDivThree(nums []int) int {
	dp := make([]int, 3)
	dp[1] = math.MinInt
	dp[2] = math.MinInt
	for _, num := range nums {
		remainder := num % 3
		dpTemp := make([]int, 3)
		copy(dpTemp, dp)
		for i := 0; i < 3; i++ {
			k := (i - remainder + 3) % 3
			dp[i] = max(dpTemp[i], dpTemp[k]+num)
		}
	}
	return dp[0]
}
