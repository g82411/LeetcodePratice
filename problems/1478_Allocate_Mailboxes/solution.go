package _478_Allocate_Mailboxes

import (
	"math"
	"sort"
)

func minDistance(houses []int, k int) int {
	n := len(houses)
	minD := make([][]int, n+1)
	dp := make([][]int, n+1)
	houses = append(houses, math.MinInt)
	sort.Ints(houses)

	for i := range minD {
		minD[i] = make([]int, n+1)
		d := make([]int, k+1)
		for j := range d {
			d[j] = math.MaxInt / 2
		}
		dp[i] = d
	}
	for i := 0; i <= n; i++ {
		for j := i; j <= n; j++ {
			for m := i; m <= j; m++ {
				minD[i][j] += abs(houses[m] - houses[(i+j)/2])
			}
		}
	}
	for i := 1; i <= n; i++ {
		dp[i][1] = minD[1][i]
	}
	for i := 1; i <= n; i++ {
		for m := 2; m <= k; m++ {
			for j := 1; j+1 <= i; j++ {
				dp[i][m] = min(dp[i][m], dp[j][m-1]+minD[j+1][i])
			}
		}
	}
	return dp[n][k]
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
