package _71_Encode_String_with_Shortest_Length

import "fmt"

// a b c d a b c d
// ^.    ^

func encode(s string) string {
	var helper func(l, r int) string
	n := len(s)
	dp := make([][]string, n)
	for i := range dp {
		dp[i] = make([]string, n)
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = string(s[i])
	}
	helper = func(l, r int) string {
		temp := s[l : r+1]
		ret := temp
		for length := 1; length < len(temp); length++ {
			if len(temp)%length > 0 {
				continue
			}
			ableUse := true
			pattern := temp[0:length]
			for i := length; i < len(temp); i += length {
				repeat := temp[i : i+length]
				ableUse = ableUse && (pattern == repeat)
				if !ableUse {
					break
				}
			}
			if ableUse {
				encodedStr := fmt.Sprintf("%v[%v]", len(temp)/length, dp[l][l+length-1])
				if len(encodedStr) < len(ret) {
					ret = encodedStr
				}
			}
		}
		return ret
	}
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			dp[i][j] = helper(i, j)
			for k := i; k < j; k++ {
				part := dp[i][k] + dp[k+1][j]
				if len(part) < len(dp[i][j]) {
					dp[i][j] = part
				}
			}
		}
	}
	return dp[0][n-1]
}
