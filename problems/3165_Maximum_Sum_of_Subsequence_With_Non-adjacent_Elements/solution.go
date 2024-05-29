package _3165_Maximum_Sum_of_Subsequence_With_Non

import "math"

// 再複習一次線段數

type Node struct {
	// 線段數的基本結構 佐子樹, 右子樹, 開端, 結束
	Left  *Node
	Right *Node
	Start int
	End   int
	// 這個題目的其中一個關鍵點，
	// Maximum Sum of Subsequence With Non-adjacent Elements
	// 這個描述本質上是House Robber
	// 但是要apply devide and concoure,
	// ret = dp00[i][j] = max{dp00[i][k]+dp00[k+1][j], dp01[i][k]+dp00[k+1][j], dp00[i][k]+dp10[k+1][j],}
	Info00 int64
	Info01 int64
	Info10 int64
	Info11 int64
}

func NewNode(start, end int, nums []int) *Node {
	ret := Node{
		Start: start,
		End:   end,
	}
	// 單點線段數的構成，只有左右都選或左右都不選兩種，剩下的沒有意義
	if start == end {
		ret.Info00 = int64(0)
		ret.Info01 = math.MinInt64 / 4
		ret.Info10 = math.MinInt64 / 4
		ret.Info11 = int64(nums[start])
		return &ret
	}
	m := (start + end) / 2
	ret.Left = NewNode(start, m, nums)
	ret.Right = NewNode(m+1, end, nums)
	// 全選的最大值 = 左右都選，左選右不選，右選左不選
	ret.Info00 = max(
		ret.Left.Info00+ret.Right.Info00,
		ret.Left.Info01+ret.Right.Info00,
		ret.Left.Info00+ret.Right.Info10,
	)
	ret.Info11 = max(
		ret.Left.Info10+ret.Right.Info01,
		ret.Left.Info11+ret.Right.Info01,
		ret.Left.Info10+ret.Right.Info11,
	)
	ret.Info10 = max(
		ret.Left.Info10+ret.Right.Info00,
		ret.Left.Info11+ret.Right.Info00,
		ret.Left.Info10+ret.Right.Info10,
	)
	ret.Info01 = max(
		ret.Left.Info00+ret.Right.Info01,
		ret.Left.Info01+ret.Right.Info01,
		ret.Left.Info00+ret.Right.Info11,
	)
	return &ret
}

func (s *Node) UpdateVal(idx int, val int) {
	// 重要 線段樹的更新機制
	if idx < s.Start || idx > s.End {
		return
	}
	if s.Start == s.End {
		s.Info00 = 0
		s.Info11 = int64(val)
		return
	}
	s.Left.UpdateVal(idx, val)
	s.Right.UpdateVal(idx, val)
	s.Info00 = max(
		s.Left.Info00+s.Right.Info00,
		s.Left.Info01+s.Right.Info00,
		s.Left.Info00+s.Right.Info10,
	)
	s.Info11 = max(
		s.Left.Info10+s.Right.Info01,
		s.Left.Info11+s.Right.Info01,
		s.Left.Info10+s.Right.Info11,
	)
	s.Info10 = max(
		s.Left.Info10+s.Right.Info00,
		s.Left.Info11+s.Right.Info00,
		s.Left.Info10+s.Right.Info10,
	)
	s.Info01 = max(
		s.Left.Info00+s.Right.Info01,
		s.Left.Info01+s.Right.Info01,
		s.Left.Info00+s.Right.Info11,
	)
}

func maximumSumSubsequence(nums []int, queries [][]int) int {
	n := len(nums)
	root := NewNode(0, n-1, nums)
	const mod = int64(1e9 + 7)
	var ret int64
	for _, q := range queries {
		root.UpdateVal(q[0], q[1])
		ret = ret + max(
			root.Info00,
			root.Info01,
			root.Info10,
			root.Info11,
		)
		ret %= mod
	}
	return int(ret)
}
