package _841_Maximum_Sum_of_Almost_Unique_Subarray

func maxSum(nums []int, m int, k int) int64 {
	occurrence := make(map[int]int)
	sum := int64(0)
	windowSum := int64(0)
	for i := 0; i < k; i++ {
		occurrence[nums[i]]++
		windowSum += int64(nums[i])
	}
	l, r := 0, k-1
	for r < len(nums) {
		uniqueNums := len(occurrence)
		if uniqueNums >= m {
			// update sum
			sum = max(sum, windowSum)
		}
		occurrence[nums[l]]--
		windowSum -= int64(nums[l])
		if occurrence[nums[l]] == 0 {
			delete(occurrence, nums[l])
		}
		l++
		r++
		if r >= len(nums) {
			break
		}
		occurrence[nums[r]]++
		windowSum += int64(nums[r])
	}
	return sum
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
