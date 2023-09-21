package _0_Regular_Expression_Matching

func isMatch(s string, p string) bool {
	// 字串比對型DP
	// dp[i][j] weather s[0:i] p[0:j] is match
	// consider case
	// s: x x x x x x x x x
	// p: y y y y y y y * .
	// if p[i] is alphabet a to z
	// dp[i][j] = s[i] == p[j] && dp[i-1][j-1]
	// if p[i] == '.'
	// dp[i][j] = dp[i-1][j-1]
	// if p[i] == '*'
	// dp[i][j] = dp[i][j - 2] || dp[i-1][j] && s[i] == p[j - 1] || p[j-1] == '.'
	// case1
	// s: x x x x x x x
	// p: a a a a z * => * can match z from 0 to n
	// case 2
	// s: x x x x x x z
	// p: a a a a a z *
	// case 3
	// s: x x x x x x z
	// p: a a a a a . *
	// dp[0][0] = true
	// dp[0][j] = true if p like y * y * y * y *
	n := len(s)
	m := len(p)
	s = "#" + s
	p = "#" + p
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for j := 2; j <= m; j++ {
		dp[0][j] = p[j] == '*' && dp[0][j-2]
	}
	isAlphabet := func(s byte) bool {
		return s-'a' >= 0 && s-'a' < 26
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if isAlphabet(p[j]) {
				dp[i][j] = s[i] == p[j] && dp[i-1][j-1]
			}
			if p[j]-'.' == 0 {
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j]-'*' == 0 {
				p1 := dp[i][j-2]
				p2 := dp[i-1][j] && (s[i] == p[j-1] || p[j-1]-'.' == 0)
				dp[i][j] = p1 || p2
			}
		}
	}
	return dp[n][m]
}
