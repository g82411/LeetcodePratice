package _1230_Toss_Strange_Coins

// dp[i][j] = 投i個硬幣 出現j個人頭的機率
// dp[i][j] = dp[i-1][j-1] * prob[i] for j <= i

func probabilityOfHeads(prob []float64, target int) float64 {
	var dp [1005][1005]float64
	n := len(prob)
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= min(i, target+1); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][0] * (1 - prob[i-1])
				continue
			}
			dp[i][j] = dp[i-1][j-1]*prob[i-1] + dp[i-1][j]*(1-prob[i-1])
		}
	}
	return dp[n][target]
}
