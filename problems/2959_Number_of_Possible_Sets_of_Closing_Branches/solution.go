package _959_Number_of_Possible_Sets_of_Closing_Branches

import "math"

func numberOfSets(n int, maxDistance int, roads [][]int) int {
	ret := 0
	for state := 0; state < (1 << n); state++ {
		dp := make([][]int, n)
		for i := range dp {
			d := make([]int, n)
			for j := range dp {
				if i != j {
					d[j] = math.MaxInt / 3
					continue
				}
				if (state >> i & 1) == 0 {
					d[j] = math.MaxInt / 3
				}
			}
			dp[i] = d
		}
		for _, road := range roads {
			src, dst, dis := road[0], road[1], road[2]
			if (state >> src & 1) == 0 {
				continue
			}
			if (state >> dst & 1) == 0 {
				continue
			}
			for i := 0; i < n; i++ {
				if (state >> i & 1) == 0 {
					continue
				}
				for j := 0; j < n; j++ {
					if (state >> j & 1) == 0 {
						continue
					}
					dp[i][j] = min(dp[i][j], dp[i][src]+dis+dp[dst][j])
					dp[i][j] = min(dp[i][j], dp[i][dst]+dis+dp[src][j])
				}
			}

		}
		isValidSet := true
		for i := 0; i < n; i++ {
			if (state >> i & 1) == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if (state >> j & 1) == 0 {
					continue
				}
				if dp[i][j] > maxDistance {
					isValidSet = false
					break
				}
			}
			if !isValidSet {
				break
			}
		}
		if isValidSet {
			ret++
		}
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
