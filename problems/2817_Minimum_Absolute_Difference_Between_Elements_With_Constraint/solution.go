package _817_Minimum_Absolute_Difference_Between_Elements_With_Constraint

import "sort"

// 我們從 i = x 開始陸續把 i - x 插入 sorted set
// 然後開始找他的 upper bound & lower bound
// 原實作透過第三方的紅黑樹來維護有序集合，但在無法存取
// 外部模組的環境下會導致編譯失敗。這裡改以 slice
// 加上 sort.SearchInts 來模擬有序集合即可。

func minAbsoluteDifference(nums []int, x int) int {
	// 使用 slice 維護已插入的元素，透過 binary search 找到
	// 與 nums[i] 最接近的前驅與後繼值
	ordered := []int{}
	ans := 1 << 30
	for i := x; i < len(nums); i++ {
		// 插入 nums[i-x]
		pos := sort.SearchInts(ordered, nums[i-x])
		ordered = append(ordered, 0)
		copy(ordered[pos+1:], ordered[pos:])
		ordered[pos] = nums[i-x]

		// 找到第一個 >= nums[i]
		pos = sort.SearchInts(ordered, nums[i])
		if pos < len(ordered) {
			ans = min(ans, ordered[pos]-nums[i])
		}
		if pos > 0 {
			ans = min(ans, nums[i]-ordered[pos-1])
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
