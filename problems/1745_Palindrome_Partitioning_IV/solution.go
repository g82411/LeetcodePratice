package _745_Palindrome_Partitioning_IV

// dp[i][j] => s[i:j] is palindrome
func checkPartitioning(s string) bool {
	n := len(s)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
		dp[i][i] = true
	}
	for i := 0; i+1 < n; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}
	for l := 3; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 2; j < n; j++ {
			if dp[0][i] && dp[i+1][j-1] && dp[j][n-1] {
				return true
			}
		}
	}
	return false
}
