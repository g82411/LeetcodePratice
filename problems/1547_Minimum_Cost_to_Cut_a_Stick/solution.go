package _547_Minimum_Cost_to_Cut_a_Stick

import (
	"math"
	"sort"
)

// 學一下區間型的bottom up
// 基本上知道子問題就是先切一刀看下一刀 但是怎麼轉移呢?
// |-----------|----|-----|------|
// i                             j
// 設ij之間存在最小切k
// 切dp[i][j]的最小切 = min{dp[i][k] + dp[k][j]} + cost for all k in (i, j)
func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0)
	cuts = append(cuts, n)
	sort.Ints(cuts)
	m := len(cuts)
	dp := make([][]int, m)
	// dp[i][j] => cuts[i:j]之間的最小切
	for i := range dp {
		d := make([]int, m)
		for j := range d {
			d[j] = math.MaxInt / 2
		}
		dp[i] = d
	}
	for i := 0; i+1 < m; i++ {
		dp[i][i+1] = 0
	}
	// 如果最小len是2那麼是不可切的
	for len := 3; len <= m; len++ {
		for i := 0; i+len-1 < m; i++ {
			j := i + len - 1
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], cuts[j]-cuts[i]+dp[i][k]+dp[k][j])
			}
		}
	}
	return dp[0][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
