package _879_Profitable_Schemes

func profitableSchemes(n int, minProfit int, groups []int, profits []int) int {
	const Mod = int64(1e9 + 7)
	K := len(groups)
	// 扯到容量就很容易背包問題
	// 1. dp[i][j][k] 完成i個任務用j個人獲得至少p profit的總數
	var dp [101][101][101]int
	dp[0][0][0] = 1
	for k := 1; k <= K; k++ {
		profit := profits[k-1]
		gSize := groups[k-1]
		for i := 0; i <= minProfit; i++ {
			for j := 0; j <= n; j++ {
				dp[k][i][j] = dp[k-1][i][j]
				// 如果使用的人數 < gSize 那人數就不夠惹拉
				if j < gSize {
					continue // QQ
				}
				dp[k][i][j] = (dp[k][i][j] + dp[k-1][max(i-profit, 0)][j-gSize]) % int(Mod)
			}
		}
	}
	sum := int64(0)
	for i := 0; i <= n; i++ {
		sum = (sum + int64(dp[K][minProfit][i])) % Mod
	}
	return int(sum)
}
