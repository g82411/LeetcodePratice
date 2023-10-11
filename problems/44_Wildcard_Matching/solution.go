package _4_Wildcard_Matching

func isMatch(s string, p string) bool {
	// 基本定義
	sLen := len(s)
	pLen := len(p)
	if s == p {
		return true
	}
	if pLen > 0 {
		allStar := true
		for i := range p {
			allStar = allStar && p[i] == '*'
		}
		if allStar {
			return true
		}
	}
	if sLen == 0 || pLen == 0 {
		return false
	}
	s = "#" + s
	p = "#" + p
	// dp [i][j] => p[0:i] match s[0:j]

	dp := make([][]bool, pLen+1)
	for i := range dp {
		dp[i] = make([]bool, sLen+1)
	}
	dp[0][0] = true
	for i := 1; i <= pLen; i++ {
		if p[i] == '*' {
			// find first match
			sIdx := 1
			for !dp[i-1][sIdx-1] && sIdx <= sLen {
				sIdx++
			}
			// if p* match s
			// p match s
			dp[i][sIdx-1] = dp[i-1][sIdx-1]
			// s: acbde....... match: acd*
			for sIdx <= sLen {
				dp[i][sIdx] = true
				sIdx++
			}
		} else if p[i] == '?' {
			// 反正?可以match any char, 所以dp[i][j] = dp[i-1][j-1]
			for j := 1; j <= sLen; j++ {
				dp[i][j] = dp[i-1][j-1]
			}
		} else {
			// 剩下的就看dp[i-1][j-1] && s[i] == p[j]
			for j := 1; j <= sLen; j++ {
				dp[i][j] = dp[i-1][j-1] && p[i] == s[j]
			}
		}
	}
	return dp[pLen][sLen]
}
