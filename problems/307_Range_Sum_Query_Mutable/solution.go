package _07_Range_Sum_Query_Mutable

type SegmentTreeNode struct {
	Value int
	Left  *SegmentTreeNode
	Right *SegmentTreeNode
	Start int
	End   int
}

func (s *SegmentTreeNode) Init(start, end int, nums []int) {
	if start == end {
		s.Value = nums[start]
		return
	}
	mid := (start + end) / 2
	if s.Left == nil {
		s.Left = &SegmentTreeNode{
			Start: start,
			End:   mid,
		}
		s.Right = &SegmentTreeNode{
			Start: mid + 1,
			End:   end,
		}
	}
	s.Left.Init(start, mid, nums)
	s.Right.Init(mid+1, end, nums)
	s.Value = s.Left.Value + s.Right.Value
}

func (s *SegmentTreeNode) UpdateSignle(idx, val int) {
	if idx < s.Start || idx > s.End {
		return
	}
	if s.Start == s.End {
		s.Value = val
		return
	}
	s.Left.UpdateSignle(idx, val)
	s.Right.UpdateSignle(idx, val)
	s.Value = s.Left.Value + s.Right.Value
}

func (s *SegmentTreeNode) QueryRange(start, end int) int {
	if end < s.Start || start > s.End {
		return 0
	}
	if start <= s.Start && end >= s.End {
		return s.Value
	}
	return s.Left.QueryRange(start, end) + s.Right.QueryRange(start, end)
}

type NumArray struct {
	root *SegmentTreeNode
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	root := &SegmentTreeNode{
		Start: 0,
		End:   n - 1,
	}
	root.Init(0, n-1, nums)
	return NumArray{
		root,
	}
}

func (this *NumArray) Update(index int, val int) {
	this.root.UpdateSignle(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.root.QueryRange(left, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
