package _935_Maximum_Strong_Pair_XOR_II

import (
	"math"
	"sort"
)

type TrieNode struct {
	Nexts [2]*TrieNode
	Count int
}

func maximumStrongPairXor(nums []int) int {
	root := &TrieNode{}
	n := len(nums)
	add := func(num int) {
		temp := root
		for i := 31; i >= 0; i-- {
			bit := (num >> i) & 1
			if temp.Nexts[bit] == nil {
				temp.Nexts[bit] = &TrieNode{}
			}
			temp = temp.Nexts[bit]
			temp.Count++
		}
	}
	remove := func(num int) {
		temp := root
		for i := 31; i >= 0; i-- {
			bit := (num >> i) & 1
			temp = temp.Nexts[bit]
			temp.Count--
		}
	}
	var dfs func(num, offset int, curr *TrieNode) int
	dfs = func(num, offset int, curr *TrieNode) int {
		if offset == -1 {
			return 0
		}
		bit := (num >> offset) & 1
		if bit == 0 {
			if curr.Nexts[1] != nil && curr.Nexts[1].Count > 0 {
				return dfs(num, offset-1, curr.Nexts[1]) + 1<<offset
			} else if curr.Nexts[0] != nil && curr.Nexts[0].Count > 0 {
				return dfs(num, offset-1, curr.Nexts[0])
			}
		} else {
			if curr.Nexts[0] != nil && curr.Nexts[0].Count > 0 {
				return dfs(num, offset-1, curr.Nexts[0]) + 1<<offset
			} else if curr.Nexts[1] != nil && curr.Nexts[1].Count > 0 {
				return dfs(num, offset-1, curr.Nexts[1])
			}
		}
		return math.MinInt / 2
	}
	// 這題的挑戰在於拆掉絕對值 跟XOR比對
	// 先將nums 排序
	// 設{x, y}為sorted pair
	// |x - y| <= min(x, y)
	// min(x, y) => x
	// x <= y <= 2 * x

	// 再來處理xor 比對, Trie一位一位處理
	sort.Ints(nums)
	l := 0
	ret := 0
	for r := 0; r < n; r++ {
		for l < n && nums[l] <= nums[r]*2 {
			add(nums[l])
			l++
		}
		ret = max(ret, dfs(nums[r], 31, root))
		remove(nums[r])
	}
	return ret
}
