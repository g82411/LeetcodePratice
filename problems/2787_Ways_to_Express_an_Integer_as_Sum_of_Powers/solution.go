package _787_Ways_to_Express_an_Integer_as_Sum_of_Powers

func numberOfWays(n int, x int) int {
	// x 最高到5 不要用fast power
	// 跟coin系列有點像
	// dp[i] 構造i的方法數
	mod := int(1e9) + 7
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		num := 1
		for t := 0; t < x; t++ {
			num *= i
		}
		// 注意這邊要由後往前掃，避免覆蓋
		for s := n; s >= num; s-- {
			dp[s] += dp[s-num]
			dp[s] %= mod
		}
	}
	return dp[n]
}
