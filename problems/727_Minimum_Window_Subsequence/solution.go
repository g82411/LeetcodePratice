package _727_Minimum_Window_Subsequence

func minWindow(s1 string, s2 string) string {
	// dp[i][j] => minimum contiguous substring part of s1[1:i], so that s2[1:j] is a subsequence of the part
	// if s1[i] == s2[i] dp[i][j] = dp[i-1][j-1] + 1
	// else dp[i][j] = dp[i-1][j] + 1 why?
	// cause we won't use s1[i] but we will consider s2[j]
	m := len(s1)
	n := len(s2)
	s1 = "#" + s1
	s2 = "#" + s2
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = math.MaxInt / 2
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				continue
			}
			dp[i][j] = dp[i-1][j] + 1
		}
	}
	l := math.MaxInt / 2
	var pos int
	for i := 1; i <= m; i++ {
		if l > dp[i][n] {
			l = dp[i][n]
			pos = i
		}
	}

	if l == math.MaxInt/2 {
		return ""
	}
	return s1[pos-l+1 : pos+1]
}
