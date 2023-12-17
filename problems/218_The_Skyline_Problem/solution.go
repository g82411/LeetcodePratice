package _18_The_Skyline_Problem

import "math"

type SegTreeNode struct {
	Left  *SegTreeNode
	Right *SegTreeNode
	Start int
	End   int
	Max   int
}

func Remove(node *SegTreeNode) {
	if node == nil {
		return
	}
	Remove(node.Left)
	Remove(node.Right)
	node = nil
	return
}

func (node *SegTreeNode) Update(start, end, val int) {
	if start >= node.End || end <= node.Start {
		return
	}
	if start <= node.Start && node.End <= end && val >= node.Max {
		Remove(node.Left)
		Remove(node.Right)
		node.Max = val
		return
	} else if val < node.Max && node.Left == nil {
		return
	}
	if node.Left == nil {
		mid := node.Start + (node.End-node.Start)/2
		node.Left = &SegTreeNode{Start: node.Start, End: mid, Max: node.Max}
		node.Right = &SegTreeNode{Start: mid, End: node.End, Max: node.Max}
	}
	node.Left.Update(start, end, val)
	node.Right.Update(start, end, val)
	node.Max = max(node.Left.Max, node.Right.Max)
}

func getSkyline(buildings [][]int) [][]int {
	var result [][]int
	var findAllStart func(node *SegTreeNode)
	root := &SegTreeNode{Start: 0, End: 1e9, Max: 0}
	for _, building := range buildings {
		root.Update(building[0], building[1], building[2])
	}
	findAllStart = func(node *SegTreeNode) {
		if node.Left == nil {
			result = append(result, []int{node.Start, node.Max})
			return
		}
		findAllStart(node.Left)
		findAllStart(node.Right)
	}
	findAllStart(root)
	if result[len(result)-1][1] > 0 {
		result = append(result, []int{math.MaxInt, 0})
	}
	var mergedResult [][]int
	for _, edge := range result {
		if len(mergedResult) > 0 && edge[1] == mergedResult[len(mergedResult)-1][1] {
			continue
		}
		mergedResult = append(mergedResult, edge)
	}
	if len(mergedResult) > 0 && mergedResult[0][1] == 0 {
		mergedResult = mergedResult[1:]
	}
	return mergedResult
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func getSkyline(buildings [][]int) [][]int {
//	// 紅黑數怎麼那麼好用
//	//
//	m := redblacktree.NewWithIntComparator()
//	// 假設有一個建築物 [[1, 5, 2], [2, 3, 8]]
//	// 1. 先把左右邊界都放進去
//	// 2. 然後把高度跟方向都放進去
//
//	for _, building := range buildings {
//		left, right, height := building[0], building[1], building[2]
//		if _, e := m.Get(left); !e {
//			m.Put(left, [][]int{})
//		}
//		if _, e := m.Get(right); !e {
//			m.Put(right, [][]int{})
//		}
//		l, _ := m.Get(left)
//		r, _ := m.Get(right)
//		m.Put(left, append(l.([][]int), []int{height, 1}))
//		m.Put(right, append(r.([][]int), []int{height, -1}))
//	}
//	set := redblacktree.NewWithIntComparator()
//	// 根據上面的東西
//	// 我們遇到1 的時候把 h 放進去
//	// 遇到 -1 的時候把 h 移除
//	// 對於每個left & right 我們只要找到當前這個set 的最大值
//	var res [][]int
//	for node := m.Iterator(); node.Next(); {
//		pos := node.Key().(int)
//		p := node.Value().([][]int)
//		for _, pair := range p {
//			height, flag := pair[0], pair[1]
//			if flag == 1 {
//				prev, found := set.Get(height)
//				if found {
//					set.Put(height, prev.(int)+1)
//				} else {
//					set.Put(height, 1)
//				}
//				continue
//			}
//			prev, _ := set.Get(height)
//			if prev.(int) == 1 {
//				set.Remove(height)
//			} else {
//				set.Put(height, prev.(int)-1)
//			}
//		}
//
//		maxHeight := 0
//		if !set.Empty() {
//			right := set.Right()
//			maxHeight = right.Key.(int)
//		}
//		if len(res) == 0 || res[len(res)-1][1] != maxHeight {
//			res = append(res, []int{pos, maxHeight})
//		}
//	}
//	return res
//}
