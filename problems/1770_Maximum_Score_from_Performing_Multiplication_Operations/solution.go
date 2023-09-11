package _770_Maximum_Score_from_Performing_Multiplication_Operations

import "math"

func maximumScore(nums []int, muls []int) int {
	n := len(nums)
	m := len(muls)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for k := 1; k <= m; k++ {
		// 左邊用了 i 個, 右邊剩下k - i個
		for i := 0; i <= k; i++ {
			j := k - i
			left := math.MinInt
			right := math.MinInt
			if i > 0 {
				left = dp[i-1][j] + nums[i-1]*muls[k-1]
			}
			if j > 0 {
				right = dp[i][j-1] + nums[n-j]*muls[k-1]
			}
			dp[i][j] = max(left, right)
		}
	}
	ans := math.MinInt
	for i := 0; i <= m; i++ {
		ans = max(ans, dp[i][m-i])
	}
	return ans
}

// func maximumScore(nums []int, multipliers []int) int {
//     n := len(nums)
//     m := len(multipliers)
//     memo := make([][]int, m+1)
//     for i := 0; i <= m; i++ {
//         d := make([]int, m+1)
//         for j := range d {
//             d[j] = math.MinInt
//         }
//         memo[i] = d
//     }
//     var dp func (l, r int) int
//     dp = func (l, r int) int {
//         nextMul := n - (r - l + 1)
//         if nextMul == m {
//             return 0
//         }
//         if memo[l][nextMul] != math.MinInt {
//             return memo[l][nextMul]
//         }
//         ans := max(
//             dp(l + 1, r) + nums[l] * multipliers[nextMul],
//             dp(l, r - 1) + nums[r] * multipliers[nextMul],
//         )
//         memo[l][nextMul] = ans
//         return ans
//     }
//     return dp(0, n - 1)

// }
// top down -> 從左往右取 取左邊 取右邊
// dp(i, k, m) = max(mul[m] * nums[i] + dp(i + 1, j, m + 1), mul[m] * nums[n - j] + dp(i, j - 1, m + 1))
// bottom up -> 從左往右取 左邊取i個 右邊取j個
// dp[i][j] => max(dp[i-1][j] + nums[i - 1] * m[i], dp[i][j - 1] + nums[n - j] * m[i])

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
