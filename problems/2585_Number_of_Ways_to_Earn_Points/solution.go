package _585_Number_of_Ways_to_Earn_Points

// opt space
// dp[j] = 達到j分的數量
// dp[j] = dp[j-k]

func waysToReachTarget(target int, q [][]int) int {
	MOD := int64(1e9) + 7
	dp := make([]int64, target+1)
	dp[0] = 1
	for _, quest := range q {
		count, score := quest[0], quest[1]
		for t := target; t >= 0; t-- {
			totalScore := 0
			for i := 0; i < count; i++ {
				totalScore += score
				if totalScore > t {
					break
				}
				dp[t] += dp[t-totalScore]
				dp[t] %= MOD
			}
		}
	}
	return int(dp[target] % MOD)
}

// 裸背包
// dp[i][j] = 做第i到題之後 能得到j分的方法
// dp[i][j] += dp[i - 1][j - k * score] k 道題
// func waysToReachTarget(target int, q [][]int) int {
//     n := len(q)
//     q = append([][]int{{0,0}}, q...)
//     dp := make([][]int64, n + 1)
//     MOD := int64(1e9 + 7)
//     for i := range dp {
//         dp[i] = make([]int64, target + 1)
//     }
//     // 一題都不做得0分的方法有一種
//     dp[0][0] = 1
//     for i := 1; i <= n; i++ {
//         for j := 0; j <= target; j++ {
//             count, score := q[i][0], q[i][1]
//             for k := 0; k <= count; k++ {
//                 if k * score > j {
//                     break
//                 }
//                 dp[i][j] = (dp[i][j] + dp[i-1][j - k * score]) % MOD
//             }
//         }
//     }
//     return int(dp[n][target] % MOD)
// }
