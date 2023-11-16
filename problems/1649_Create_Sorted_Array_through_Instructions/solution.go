package _649_Create_Sorted_Array_through_Instructions

import "sort"

type SegTreeNode struct {
	Left  *SegTreeNode
	Right *SegTreeNode
	Start int
	End   int
	Freq  int
}

func initSegTree(root *SegTreeNode, start, end int) {
	if start == end {
		root.Freq = 0
		return
	}
	mid := (start + end) / 2
	if root.Left == nil {
		root.Left = &SegTreeNode{
			Start: start,
			End:   mid,
		}
		root.Right = &SegTreeNode{
			Start: mid + 1,
			End:   end,
		}
	}
	initSegTree(root.Left, start, mid)
	initSegTree(root.Right, mid+1, end)
	root.Freq = root.Left.Freq + root.Right.Freq
}

func increaseFreq(root *SegTreeNode, id int) {
	if id < root.Start || id > root.End {
		return
	}
	if root.Start == root.End {
		root.Freq++
		return
	}
	increaseFreq(root.Left, id)
	increaseFreq(root.Right, id)
	root.Freq = root.Left.Freq + root.Right.Freq
}

func queryRange(root *SegTreeNode, start, end int) int {
	if end < root.Start || start > root.End {
		return 0
	}
	if start <= root.Start && end >= root.End {
		return root.Freq
	}
	return queryRange(root.Left, start, end) + queryRange(root.Right, start, end)
}

func createSortedArray(instructions []int) int {
	const M = int64(1e9 + 7)
	temp := make([]int, len(instructions))
	copy(temp, instructions)
	sort.Ints(temp)
	set := make(map[int]int)
	num2idx := make(map[int]int)
	// 離散化, 只關心數值 不關心相加結果
	// https://www.youtube.com/watch?v=rNdv9xlP8rk
	i := 0
	for _, num := range temp {
		if _, e := set[num]; e {
			continue
		}
		set[num]++
		num2idx[num] = i
		i++
	}
	n := len(num2idx)
	root := &SegTreeNode{
		Start: 0,
		End:   n - 1,
	}
	initSegTree(root, 0, n-1)
	ret := int64(0)
	for _, x := range instructions {
		left := queryRange(root, 0, num2idx[x]-1)
		right := queryRange(root, num2idx[x]+1, n-1)
		ret += int64(min(left, right))
		ret %= M
		increaseFreq(root, num2idx[x])
	}
	return int(ret)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
