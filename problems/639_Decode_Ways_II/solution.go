package _39_Decode_Ways_II

// X X X X 1 *
// dp[i] = dp[i-1] * 9 + dp[i-2] * 9
// X X X X * *
// dp[i] = dp[i-1] * 9 + dp[i-2] * 9 + dp[i-2] *

func isDigit(c byte) bool {
	if c-'0' > 9 {
		return false
	}
	if c-'0' < 0 {
		return false
	}
	return true
}

func numDecodings(s string) int {
	n := len(s)
	s = "#" + s
	const M = int(1e9) + 7
	dp := make([]int, n+1)
	dp[0] = 1
	// dp[1] = ?
	if s[1] == '*' {
		dp[1] = 9
	} else if s[1] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	for i := 2; i <= n; i++ {
		if isDigit(s[i]) {
			// if XXXX[1~9]
			dp[i] = dp[i-1]
			if s[i] == '0' {
				dp[i] = 0
			}
			if isDigit(s[i-1]) {
				num := (s[i-1]-'0')*10 + s[i] - '0'
				if num >= 10 && num <= 26 {
					dp[i] = (dp[i] + dp[i-2]) % M
				}
			} else {
				// XXXX*[1~9]
				// if s[i] 0~6 -> * will be 1 or 2
				if s[i] <= '6' {
					dp[i] = (dp[i] + dp[i-2]*2) % M
				} else {
					dp[i] = (dp[i] + dp[i-2]) % M
				}
			}
		} else {
			dp[i] = dp[i-1] * 9 % M
			if s[i-1] == '1' {
				dp[i] = (dp[i] + dp[i-2]*9) % M
			}
			if s[i-1] == '2' {
				dp[i] = (dp[i] + dp[i-2]*6) % M
			}
			if s[i-1] == '*' {
				dp[i] = (dp[i] + dp[i-2]*15) % M
			}
		}
	}
	return dp[n]
}
