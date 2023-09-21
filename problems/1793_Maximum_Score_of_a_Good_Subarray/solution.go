package _793_Maximum_Score_of_a_Good_Subarray

import "math"

func maximumScore(nums []int, k int) int {
	// 基本的雙指針
	n := len(nums)
	l, r := k, k
	minVal := nums[k]
	ret := 0
	for l >= 0 || r < n {
		for l >= 0 && nums[l] >= minVal {
			l--
		}
		for r < n && nums[r] >= minVal {
			r++
		}
		// 這題唯一的難度, 注意當上面兩個迴圈結束後, r & l 都在界外, 長度要扣掉2
		ret = max(ret, (r-l-1)*minVal)
		lmin := math.MinInt
		rmin := math.MinInt
		if l >= 0 {
			lmin = nums[l]
		}
		if r < n {
			rmin = nums[r]
		}
		minVal = max(lmin, rmin)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
