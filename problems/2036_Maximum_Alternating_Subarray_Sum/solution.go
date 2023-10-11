package _036_Maximum_Alternating_Subarray_Sum

import "math"

func maximumAlternatingSubarraySum(nums []int) int64 {
	ret := int64(math.MinInt / 2)
	sum1 := int64(math.MinInt / 2)
	sum2 := int64(0)
	for _, num := range nums {
		s1T := sum1
		s2T := sum2
		sum1 = max(s2T+int64(num), int64(num))
		sum2 = s1T - int64(num)
		ret = max(ret, sum1)
		ret = max(ret, sum2)
	}
	return ret
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
