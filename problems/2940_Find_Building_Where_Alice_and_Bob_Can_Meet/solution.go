package _940_Find_Building_Where_Alice_and_Bob_Can_Meet

// in query[i] {a, b} 找到第一個heights[k] > max(heights[a], heights[b])
// 所以從最右邊開始處理 然後把處理過的heights 放到有序容器裡面

// in query[i] {a, b} 找到第一個heights[k] > max(heights[a], heights[b])
// 所以從最右邊開始處理 然後把處理過的heights 放到有序容器裡面

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"sort"
)

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	// 先確保for all {a, b} in query[i] a < b
	for i := range queries {
		sort.Ints(queries[i])
		// 保留原本的Idx 因為後面要排序
		queries[i] = append(queries[i], i)
	}
	// 最右邊的先處理
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][1] > queries[j][1]
	})
	// 有序集合怎摸辦, 直接送一顆紅黑樹
	rbt := redblacktree.NewWithIntComparator()
	rets := make([]int, len(queries))
	i := len(heights) - 1
	//
	for _, query := range queries {
		a, b, idx := query[0], query[1], query[2]
		// 將在b右邊的東西都擺進有序集合裡面惹
		for i >= b {
			for !rbt.Empty() && heights[i] >= rbt.Left().Key.(int) {
				rbt.Remove(rbt.Left().Key)
			}
			rbt.Put(heights[i], i)
			i--
		}
		if heights[a] < heights[b] || a == b {
			rets[idx] = b
			continue
		}
		target := max(heights[a], heights[b])
		f, found := rbt.Ceiling(target + 1)
		if found {
			rets[idx] = f.Value.(int)
			continue
		}
		rets[idx] = -1
		// 如果找到的元素等於目標值，則尋找下一個元素

	}
	return rets
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
