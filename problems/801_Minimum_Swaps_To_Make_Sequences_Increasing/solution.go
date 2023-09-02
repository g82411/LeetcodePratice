package _01_Minimum_Swaps_To_Make_Sequences_Increasing

import "math"

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	swap := make([]int, n)
	keep := make([]int, n)
	for i := range swap {
		swap[i] = math.MaxInt
		keep[i] = math.MaxInt
	}
	swap[0] = 1
	keep[0] = 0
	for i := 1; i < n; i++ {
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			// no need to swap
			keep[i] = keep[i-1]
			// prev step swap, so next step should swap to keep sequence increasing
			swap[i] = swap[i-1] + 1
		}
		// notice that need swap != can swap, so need to check
		if nums1[i-1] < nums2[i] && nums2[i-1] < nums1[i] {
			// we don't swap i but swap i - 1
			keep[i] = min(swap[i-1], keep[i])
			swap[i] = min(swap[i], keep[i-1]+1)
		}
	}
	return min(keep[n-1], swap[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
