package _10_Split_Array_Largest_Sum

import "math"

// func splitArray(nums []int, k int) int {
//     n := len(nums)
//     prefixSum := make([]int, n + 1)
//     left, right := 0, 0
//     for i, num := range nums {
//         prefixSum[i + 1] = prefixSum[i] + num
//         left = max(left, num)
//     }
//     right = prefixSum[n]
//     canSplit := func(targetSum int) bool {
//         cnt := 1
//         curSum := 0
//         for i := 0; i < n; i ++ {
//             curSum += nums[i];
//             if (curSum > targetSum) {
//                 curSum = nums[i]
//                 cnt++
//                 if (cnt > k) {
//                     return false
//                 }
//             }
//         }
//         return true
//     }
//     for left < right {
//         targetSum := (left + right) / 2
//         if canSplit(targetSum) {
//             right = targetSum
//             continue
//         }
//         left = targetSum + 1
//     }
//     return right

// }

// dp[i][j] => min max sum for split nums[0:i] into k parts
// dp[i][j] = min(dp[i][j], max(dp[j-1][k-1], sum(nums[i:j])))
func splitArray(nums []int, k int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		d := make([]int, k+1)
		for j := range d {
			d[j] = math.MaxInt / 2
		}
		dp[i] = d
	}
	dp[0][0] = 0
	nums = append([]int{0}, nums...)
	for i := 1; i <= n; i++ {
		for p := 1; p <= min(i, k); p++ {
			suffixSum := 0
			for j := i; j >= p; j-- {
				suffixSum += nums[j]
				dp[i][p] = min(dp[i][p], max(dp[j-1][p-1], suffixSum))
			}
		}
	}
	return dp[n][k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
