package _18_The_Skyline_Problem

import "github.com/emirpasic/gods/trees/redblacktree"

func getSkyline(buildings [][]int) [][]int {
	// 紅黑數怎麼那麼好用
	//
	m := redblacktree.NewWithIntComparator()
	// 假設有一個建築物 [[1, 5, 2], [2, 3, 8]]
	// 1. 先把左右邊界都放進去
	// 2. 然後把高度跟方向都放進去

	for _, building := range buildings {
		left, right, height := building[0], building[1], building[2]
		if _, e := m.Get(left); !e {
			m.Put(left, [][]int{})
		}
		if _, e := m.Get(right); !e {
			m.Put(right, [][]int{})
		}
		l, _ := m.Get(left)
		r, _ := m.Get(right)
		m.Put(left, append(l.([][]int), []int{height, 1}))
		m.Put(right, append(r.([][]int), []int{height, -1}))
	}
	set := redblacktree.NewWithIntComparator()
	// 根據上面的東西
	// 我們遇到1 的時候把 h 放進去
	// 遇到 -1 的時候把 h 移除
	// 對於每個left & right 我們只要找到當前這個set 的最大值
	var res [][]int
	for node := m.Iterator(); node.Next(); {
		pos := node.Key().(int)
		p := node.Value().([][]int)
		for _, pair := range p {
			height, flag := pair[0], pair[1]
			if flag == 1 {
				prev, found := set.Get(height)
				if found {
					set.Put(height, prev.(int)+1)
				} else {
					set.Put(height, 1)
				}
				continue
			}
			prev, _ := set.Get(height)
			if prev.(int) == 1 {
				set.Remove(height)
			} else {
				set.Put(height, prev.(int)-1)
			}
		}

		maxHeight := 0
		if !set.Empty() {
			right := set.Right()
			maxHeight = right.Key.(int)
		}
		if len(res) == 0 || res[len(res)-1][1] != maxHeight {
			res = append(res, []int{pos, maxHeight})
		}
	}
	return res
}
