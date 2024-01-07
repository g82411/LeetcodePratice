package _837_New_21_Game

// 這副牌共[1, maxPts]種點數
// Alice在拿到k點後停牌
// 問Alice拿到不大於n點的機率
// 其中 n >= k
// n k m <= 1e4
func new21Game(n int, k int, maxPts int) float64 {
	if maxPts+k-1 <= n || k == 0 {
		// 如果最大值 + 停牌 - 1 <= n 意味著拿到k-1之後不停牌一定會>n
		return 1.0
	}
	// dp[i] 拿到i 點的機率
	// dp[i] = dp[i-k] + dp[i-(k+1)] +... +dp[i-maxPts] / maxPts
	dp := make([]float64, n+1)
	dp[0] = 1.0
	sumOfProb := 1.0
	res := 0.0
	for i := 1; i <= n; i++ {
		dp[i] = sumOfProb / float64(maxPts)
		if i >= k {
			// 可以停牌的狀況
			res += dp[i]
		}
		if i < k {
			// 不能停牌 機率相加
			sumOfProb += dp[i]
		}
		// i 超過maxPts, 意味著不能從i - maxPts拿到我要的值
		if i >= maxPts {
			sumOfProb -= dp[i-maxPts]
		}
	}
	return res

}

// DFS TLE
// func new21Game(n int, k int, maxPts int) float64 {
//     var dfs func (score int) float64
//     var cache [30000]float64
//     for i := range cache {
//         cache[i] = -1
//     }
//     dfs = func (score int) float64 {
//         if score >= k {
//             if score <= n {
//                 return 1.0
//             } else {
//                 return 0.0
//             }
//         }
//         if cache[score] >= 0 {
//             return cache[score]
//         }
//         p := 0.0
//         for i := 1; i <= maxPts; i++ {
//             p += dfs(score + i)
//         }
//         p /= float64(maxPts)
//         cache[score] = p
//         return cache[score]
//     }
//     return dfs(0)
// }
