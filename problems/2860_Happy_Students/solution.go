package _860_Happy_Students

import "sort"

func countWays(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ret := 0
	// x x x x x i | i + 1 x x x
	// 選i + 1 => nums[i + 1] > i + 1
	for i := 0; i+1 < n; i++ {
		if i+1 > nums[i] && i+1 < nums[i+1] {
			ret++
		}
	}
	if 0 < nums[0] {
		ret++
	}

	if n > nums[n-1] {
		ret++
	}
	return ret
}
