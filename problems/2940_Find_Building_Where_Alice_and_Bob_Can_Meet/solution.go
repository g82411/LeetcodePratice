package _940_Find_Building_Where_Alice_and_Bob_Can_Meet

// in query[i] {a, b} 找到第一個heights[k] > max(heights[a], heights[b])
// 所以從最右邊開始處理 然後把處理過的heights 放到有序容器裡面

// in query[i] {a, b} 找到第一個heights[k] > max(heights[a], heights[b])
// 所以從最右邊開始處理 然後把處理過的heights 放到有序容器裡面

import "sort"

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
	rets := make([]int, len(queries))
	stack := []int{}
	i := len(heights) - 1
	for _, query := range queries {
		a, b, idx := query[0], query[1], query[2]
		// 將在 b 右邊的元素加入單調遞減 stack
		for i >= b {
			for len(stack) > 0 && heights[i] >= heights[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
			i--
		}
		if heights[a] < heights[b] || a == b {
			rets[idx] = b
			continue
		}
		target := max(heights[a], heights[b])
		ansIdx := -1
		for j := len(stack) - 1; j >= 0; j-- {
			if stack[j] > b && heights[stack[j]] > target {
				ansIdx = stack[j]
				break
			}
		}
		rets[idx] = ansIdx
	}
	return rets
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
