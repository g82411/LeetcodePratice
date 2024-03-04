package _1278_Palindrome_Partitioning_III

import "math"

func palindromePartition(s string, K int) int {
	// 建立一個count數組計算從s[i:j]要變化多少個才會是回文
	// 這邊厲害的是他是一個II行DP
	m := len(s)
	s = "#" + s
	count := make([][]int, m+1)
	for i := range count {
		count[i] = make([]int, m+1)
	}
	for l := 2; l <= m; l++ {
		for i := 1; i+l-1 <= m; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				count[i][j] = count[i+1][j-1]
				continue
			}
			count[i][j] = count[i+1][j-1] + 1
		}
	}
	// dp[i][k] => 將s[1:i] 分成k份所需要的 最小的回文數量
	dp := make([][]int, m+1)
	for i := range dp {
		d := make([]int, K+1)
		for j := range d {
			d[j] = math.MaxInt / 2
		}
		dp[i] = d
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		// 注意 if k = 3 i = 2 s[1:i]最多只能分成k份
		for k := 1; k <= min(i, K); k++ {
			dp[i][k] = math.MaxInt / 2
			for j := k; j <= i; j++ {
				// s[1:j-1]分成k-1份
				dp[i][k] = min(dp[i][k], dp[j-1][k-1]+count[j][i])
			}
		}
	}
	return dp[m][K]
}
