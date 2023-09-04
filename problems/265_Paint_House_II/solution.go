package _65_Paint_House_II

func minCostII(costs [][]int) int {
	// 有趣的題目 問題點在於怎麼找到前一個的最小值
	// 解法是前一個只有兩種 最小的 跟第二小的 兩種顏色不一樣
	// 如果下一輪跟最小的顏色一樣 就換第二小的
	houseAmount := len(costs)
	colorAmount := len(costs[0])
	dp := make([][]int, houseAmount)
	for i := range dp {
		dp[i] = make([]int, colorAmount)
	}
	copy(dp, costs)
	min1, min2 := -1, -1
	for i := 0; i < houseAmount; i++ {
		last1, last2 := min1, min2
		min1, min2 = -1, -1
		for j := 0; j < colorAmount; j++ {
			if j != last1 && last1 >= 0 {
				dp[i][j] += dp[i-1][last1]
			} else if last2 >= 0 {
				dp[i][j] += dp[i-1][last2]
			}
			if min1 < 0 || dp[i][j] < dp[i][min1] {
				min2 = min1
				min1 = j
			} else if min2 < 0 || dp[i][j] < dp[i][min2] {
				min2 = j
			}
		}
	}
	return dp[houseAmount-1][min1]
}
