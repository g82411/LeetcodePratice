package _473_Paint_House_III

import "math"

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	memo := make(map[[3]int]int)
	var dfs func(prevColor, neighborhoodCount, curr int) int
	dfs = func(pC, nC, curr int) int {
		if curr == m {
			if nC == target {
				return 0
			}
			return math.MaxInt / 2
		}
		if nC > target {
			return math.MaxInt / 2
		}
		v, e := memo[[3]int{pC, nC, curr}]
		if e {
			return v
		}
		minCost := math.MaxInt / 2
		if houses[curr] > 0 {
			newNC := nC
			if pC != houses[curr] {
				newNC++
			}
			minCost = dfs(houses[curr], newNC, curr+1)
		} else {
			currCosts := cost[curr]
			for i, curCost := range currCosts {
				paint := i + 1
				newNC := nC
				if pC != paint {
					newNC++
				}
				minCost = min(minCost, curCost+dfs(paint, newNC, curr+1))
			}
		}
		memo[[3]int{pC, nC, curr}] = minCost
		return minCost

	}
	ans := dfs(0, 0, 0)
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}

// func minCost(houses []int, cost [][]int, m int, n int, target int) int {
//     // dp[k][p][i] -> k組房子, 第p個房子塗上第i個顏色
//     // dp[k][p][i] = min{dp[k-1][p-1][i] + cost[p][i], dp[k][p-1][i]}
//     if n == 1 && houses[0] > 0 {
//       return 0
//     } else if n == 1 && houses[0] == 0 {
//       return cost[0][0]
//     }
//     dp := make([][][]int, target + 1)
//     for i := range dp {
//         dp[i] = make([][]int, m + 1)
//         for j := range dp[i] {
//             d := make([]int, n + 1)
//             for k := range d {
//                 d[k] = math.MaxInt / 2
//             }
//             dp[i][j] = d
//         }
//     }
//     for i := range dp[0][0] {
//         dp[0][0][i] = 0
//     }

//     for k := 1; k <= target; k++ {
//         for i := k; i<= m; i++ {
//             colorI := houses[i-1]
//             colorJ := 0
//             if i > 1 {
//                 colorJ = houses[i-2]
//             }
//             sI, eI := colorI, colorI
//             sJ, eJ := colorJ, colorJ
//             if colorI == 0 {
//                 sI, eI = 1, n
//             }
//             if colorJ == 0 {
//                 sJ, eJ = 1, n
//             }

//             for ci := sI; ci <= eI; ci++ {
//                 v := 0
//                 if ci != colorI {
//                     v = cost[i-1][ci-1]
//                 }
//                 for cj := sJ; cj <= eJ; cj++ {
//                     if ci == cj {
//                         dp[k][i][ci] = min(dp[k][i][ci], dp[k][i-1][cj] + v)
//                         continue
//                     }
//                     // 多分一組
//                     dp[k][i][ci] = min(dp[k][i][ci], dp[k-1][i-1][cj] + v)
//                 }
//             }
//         }
//     }
//     ret := math.MaxInt
//     for i := 0; i <= n; i++ {
//         ret = min(ret, dp[target][m][i])
//     }
//     if ret == math.MaxInt / 2 {
//         return -1
//     }
//     return ret
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
