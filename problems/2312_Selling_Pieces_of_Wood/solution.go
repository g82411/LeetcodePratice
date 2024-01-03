package _2312_Selling_Pieces_of_Wood

// dp[i][j] => the max money we can earn after cutting an i * j pieces wood
// dp[i][j] => max(max(dp[k][j] + dp[i][j] for k in range i), dp[i][c] + dp[i][j-c] for c in ragne j) for k in
// 總而言之 砍一刀
func sellingWood(m int, n int, prices [][]int) int64 {
	var dp [205][205]int64
	for _, price := range prices {
		r, c, p := price[0], price[1], price[2]
		dp[r][c] = int64(p)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k < i; k++ {
				dp[i][j] = max(dp[i][j], dp[k][j]+dp[i-k][j])
			}
			for k := 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[i][j-k])
			}
		}
	}
	return dp[m][n]
}
