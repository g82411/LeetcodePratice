package _90_Domino_and_Tromino_Tiling

func numTilings(n int) int {
	// 0 完整填滿
	// 1 上或下突一塊

	dp := make([][2]int, n+1)
	dp[0][0] = 1
	dp[1][0] = 1
	mod := int(1e9) + 7
	for i := 2; i <= n; i++ {
		// complete cover i, should have dp[i - 1][0] + dp[i-2][0] (add one domino vertical or add two domino horizontal)
		// + dp[i - 2][1] * 2 (Tromino which make row 1 or row 2)
		dp[i][0] = (dp[i-1][0]%mod+dp[i-2][0]%mod)%mod + (dp[i-1][1]*2)%mod
		dp[i][1] = (dp[i-2][0] + dp[i-1][1]) % mod
	}
	return dp[n][0]
}
