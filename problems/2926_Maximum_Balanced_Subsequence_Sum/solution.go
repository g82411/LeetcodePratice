package _926_Maximum_Balanced_Subsequence_Sum

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"math"
)

func maxBalancedSubsequenceSum(nums []int) int64 {
	n := len(nums)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nums[i] - i
	}

	rbt := redblacktree.NewWithIntComparator()
	var ret int64 = math.MinInt64

	for i := 0; i < n; i++ {
		x := arr[i]
		node, found := rbt.Floor(x)
		if found {
			rbt.Put(x, max(int64(nums[i]), node.Value.(int64)+int64(nums[i])))
		} else {
			rbt.Put(x, int64(nums[i]))
		}
		temp, _ := rbt.Get(x)
		if ret < temp.(int64) {
			ret = temp.(int64)
		}

		// 使用迭代器來移除小於或等於當前值的元素
		it := rbt.IteratorAt(rbt.GetNode(x))

		var waitToRemove []int
		for it.Next() {
			if it.Value().(int64) <= temp.(int64) {
				waitToRemove = append(waitToRemove, it.Key().(int))
			} else {
				break
			}
		}
		for _, v := range waitToRemove {
			rbt.Remove(v)
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
