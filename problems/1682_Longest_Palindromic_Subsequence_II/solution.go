package _682_Longest_Palindromic_Subsequence_II

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][][26]int, n+1)
	for i := range dp {
		dp[i] = make([][26]int, n+1)
	}
	for i := 0; i+1 < n; i++ {
		c1 := s[i] - 'a'
		c2 := s[i+1] - 'a'
		if c1 == c2 {
			dp[i][i+1][c1] = 2
		}
	}
	for l := 3; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			c1 := s[i] - 'a'
			c2 := s[j] - 'a'
			if c1 == c2 {
				for c := 0; c < 26; c++ {
					if byte(c) != c1 {
						dp[i][j][c1] = max(dp[i][j][c1], dp[i+1][j-1][c]+2)
					}
				}
				for c := 0; c < 26; c++ {
					if byte(c) != c1 {
						dp[i][j][c] = dp[i+1][j-1][c]
					}
				}
			} else {
				dp[i][j][c1] = dp[i][j-1][c1]
				dp[i][j][c2] = dp[i+1][j][c2]
				for c := 0; c < 26; c++ {
					if byte(c) != c1 && byte(c) != c2 {
						dp[i][j][c] = dp[i+1][j-1][c]
					}
				}
			}
		}
	}
	ret := 0
	for c := 0; c < 26; c++ {
		ret = max(ret, dp[0][n-1][c])
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
