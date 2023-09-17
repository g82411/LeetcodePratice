package _64_Strange_Printer

import "math"

// 這題很賤的地方就是透過很短的字串讓你看不到狀態轉移跟子問題
// 然後子問題也不好想
// 其實是可以根據字串中的相同char去分割
// 考慮
// x x x x a x x x x x x a x x x x x x a x x x
//         i             k             j
// 就可以理解決 dp[i][j] = dp[i][k-1] + dp[i][k + 1]
// x x x x a x x x a x x a x x x a x x a x x x
//         i       m     k       n     j
// 考慮為
// dp[i][j] = min(
//    dp[i][k-1] + dp[i][k + 1],
//    dp[i][m-1] + dp[i][m + 1],
//    dp[i][n-1] + dp[i][n + 1],
// )
//
//

func strangePrinter(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		d := make([]int, n)
		for j := range d {
			d[j] = math.MaxInt / 2
		}
		dp[i] = d
	}
	// 長度為1的字串只需要把自己印出來
	for i := range dp {
		dp[i][i] = 1
	}
	// 從區間大小2開始推進
	for len := 2; len <= n; len++ {
		for i := 0; i <= n-len; i++ {
			j := i + len - 1
			dp[i][j] = 1 + dp[i+1][j]
			for k := i + 1; k <= j; k++ {
				if s[i] == s[k] {
					if k+1 > j {
						dp[i][j] = min(dp[i][j], dp[i][k-1]+0)
						continue
					}
					dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
