package _926_Maximum_Balanced_Subsequence_Sum

import (
	"math"
	"sort"
)

func maxBalancedSubsequenceSum(nums []int) int64 {
	n := len(nums)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nums[i] - i
	}

	// pairs 以 arr 值遞增排序
	type pair struct {
		key int
		val int64
	}
	var pairs []pair
	var ret int64 = math.MinInt64

	for i := 0; i < n; i++ {
		x := arr[i]
		// 找到 <= x 的最後一個元素
		pos := sort.Search(len(pairs), func(j int) bool { return pairs[j].key > x })
		cur := int64(nums[i])
		if pos > 0 {
			cur = max(cur, pairs[pos-1].val+int64(nums[i]))
		}
		// 插入或更新
		if pos < len(pairs) && pairs[pos].key == x {
			if cur > pairs[pos].val {
				pairs[pos].val = cur
			}
		} else {
			pairs = append(pairs, pair{})
			copy(pairs[pos+1:], pairs[pos:])
			pairs[pos] = pair{key: x, val: cur}
		}
		// 移除後續 val 小於等於當前值的元素，保持遞減
		idx := pos + 1
		for idx < len(pairs) && pairs[idx].val <= cur {
			idx++
		}
		if idx > pos+1 {
			copy(pairs[pos+1:], pairs[idx:])
			pairs = pairs[:len(pairs)-(idx-(pos+1))]
		}

		if ret < cur {
			ret = cur
		}
	}

	return ret
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
