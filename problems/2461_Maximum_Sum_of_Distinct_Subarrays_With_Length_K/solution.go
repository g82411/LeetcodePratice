package _461_Maximum_Sum_of_Distinct_Subarrays_With_Length_K

func maximumSubarraySum(nums []int, k int) int64 {
	occurrency := make(map[int]int)
	l := 0
	sum := int64(0)
	maxSum := int64(0)
	for r, num := range nums {
		occurrency[num]++
		sum += int64(num)
		if r-l+1 < k {
			continue
		}
		if len(occurrency) == k {
			maxSum = max(maxSum, sum)
		}
		sum -= int64(nums[l])
		occurrency[nums[l]]--
		if occurrency[nums[l]] <= 0 {
			delete(occurrency, nums[l])
		}
		l++

	}
	return maxSum
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
