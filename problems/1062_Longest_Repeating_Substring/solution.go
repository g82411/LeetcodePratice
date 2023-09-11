package _062_Longest_Repeating_Substring

func longestRepeatingSubstring(s string) int {
	n := len(s)
	s = "#" + s
	// 當作LCS來解 但是要記得可以用兩個字串做比較 是很常見的DP作法
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i != j && s[i] == s[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
