package _97_Interleaving_String

func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	o := len(s3)
	if n+m != o {
		return false
	}
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			}
			comparsion := s3[i+j-1]
			if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == comparsion
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == comparsion
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == comparsion) || (dp[i][j-1] && s2[j-1] == comparsion)
			}
		}
	}
	return dp[n][m]
}
