package _915_Length_of_the_Longest_Subsequence_That_Sums_to_Target

func lengthOfLongestSubsequence(nums []int, target int) int {
	// 0 1 背包問題
	dp := make([]int, target+1)
	for _, num := range nums {
		for t := target; t >= num; t-- {
			if t == num || dp[t-num] > 0 {
				dp[t] = max(dp[t], 1+dp[t-num])
			}
		}
	}
	if dp[target] > 0 {
		return dp[target]
	}
	return -1
}
