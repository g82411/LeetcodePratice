package _2916_Subarrays_Distinct_Element_Sum_of_Squares_II

//設nums = X X X X X X k X X X X i
//dp[i] = 以nums[i]為結尾, 所有subarray 的distincts nums ^ 2
//nums[k] == nums[i] 所以對於k之後到i-1之前 distincts nums部會因為i而變化
//並且square[k:i-1] 也應該等於square[k:i] (因為i == k)
//dp[i]
//= sum{square[j:i]} for j in [0, i]
//= sum{square[j:i]} for j in [0, k] + sum{square[j:i]} for j in [k+1, i]
//= sum{square[j:i-1]} for j in [0, k]
// +sum{square[j:i]} for j in [k+1, i]
//
//= sum{square[j:i-1]} for j in [0, k]
// +sum{(count[j:i-1] + 1) ^ 2} for j in [k+1, i-1]
//
//= sum{square[j:i-1]} for j in [0, k]
// +sum{count[j:i-1] ^ 2 + 2 * count[j:i-1] + 1} for j in [k+1,  i-1]
//
//= sum{square[j:i-1]} for j in [0, k]
// +sum{square[j:i-1] + 2 * count[j:i-1] + 1} for j in [k+1, i i-1]
//
//= sum{square[j:i-1]} for j in [0, k] + sum{square[j:i-1]} for j in [k+1,  i-1]+
//  sum{ 2 * count[j:i-1] + 1} for j in [k+1,  i-1] + 1
//
//= dp[i-1] +  2 * sum{count[j:i-1]} + i - k for j in [k+1, i-1] + 1

type SegTreeNode struct {
	Left  *SegTreeNode
	Right *SegTreeNode
	Start int
	End   int
	Val   int64
	Delta int64
	Tag   bool
}

func NewSegTreeNodeBySlice(rStart, rEnd int, nums []int) *SegTreeNode {
	ret := SegTreeNode{
		Tag:   false,
		Start: rStart,
		End:   rEnd,
	}
	if rStart == rEnd {
		ret.Val = int64(nums[rStart])
		return &ret
	}
	mid := (rStart + rEnd) / 2
	if ret.Left == nil {
		ret.Left = NewSegTreeNodeBySlice(rStart, mid, nums)
		ret.Right = NewSegTreeNodeBySlice(mid+1, rEnd, nums)
		ret.Val = ret.Left.Val + ret.Right.Val
	}
	return &ret
}

func NewSegTreeNodeByVal(rStart, rEnd int, val int) *SegTreeNode {
	ret := SegTreeNode{
		Tag:   false,
		Start: rStart,
		End:   rEnd,
	}
	if rStart == rEnd {
		ret.Val = int64(val)
		return &ret
	}
	mid := (rStart + rEnd) / 2
	if ret.Left == nil {
		ret.Left = NewSegTreeNodeByVal(rStart, mid, val)
		ret.Right = NewSegTreeNodeByVal(mid+1, rEnd, val)
		ret.Val = ret.Left.Val + ret.Right.Val
	}
	return &ret
}

func (node *SegTreeNode) pushDown() {
	if !node.Tag {
		return
	}
	if node.Left == nil {
		return
	}
	node.Left.Val += node.Delta * int64(node.Left.End-node.Left.Start+1)
	node.Left.Delta += node.Delta
	node.Left.Tag = true

	node.Right.Val += node.Delta * int64(node.Right.End-node.Right.Start+1)
	node.Right.Delta += node.Delta
	node.Right.Tag = true

	node.Tag = false
	node.Delta = int64(0)
}

func (node *SegTreeNode) UpdateRangeBy(rStart, rEnd int, val int64) {
	if rEnd < node.Start || rStart > node.End {
		return
	}
	if rStart <= node.Start && node.End <= rEnd {
		node.Val += val * int64(node.End-node.Start+1)
		node.Delta += val
		node.Tag = true
		return
	}
	if node.Left == nil {
		return
	}
	node.pushDown()
	node.Left.UpdateRangeBy(rStart, rEnd, node.Delta+val)
	node.Right.UpdateRangeBy(rStart, rEnd, node.Delta+val)
	node.Delta = int64(0)
	node.Tag = false
	node.Val = node.Left.Val + node.Right.Val
}

func (node *SegTreeNode) QueryRange(rStart, rEnd int) int64 {
	if rEnd < node.Start || rStart > node.End {
		return int64(0)
	}
	if rStart <= node.Start && node.End <= rEnd {
		return node.Val
	}
	if node.Left == nil {
		return node.Val
	}
	node.pushDown()
	ret := node.Left.QueryRange(rStart, rEnd) + node.Right.QueryRange(rStart, rEnd)
	node.Val = node.Left.Val + node.Right.Val
	return ret
}

func sumCounts(nums []int) int {
	const Mod = int(1e9 + 7)
	n := len(nums)
	posMap := make(map[int]int)
	prev := make([]int, n)
	for i := range prev {
		prev[i] = -1
	}

	for i, val := range nums {
		if _, e := posMap[val]; e {
			prev[i] = posMap[val]
		}
		posMap[val] = i
	}

	root := NewSegTreeNodeByVal(0, n-1, 0)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		k := prev[i]
		dp[i] = 0
		if i > 0 {
			dp[i] = dp[i-1]
		}
		dp[i] += 2*int(root.QueryRange(k+1, i-1)) + i - 1 - k
		dp[i] += 1
		dp[i] %= Mod
		// k + 1 到i之間 產生了一個新的distinct
		root.UpdateRangeBy(k+1, i, int64(1))
	}
	ret := int64(0)
	for _, val := range dp {
		ret = (ret + int64(val)) % int64(Mod)
	}

	return int(ret)
}
