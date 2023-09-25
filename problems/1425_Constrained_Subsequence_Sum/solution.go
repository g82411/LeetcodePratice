package _425_Constrained_Subsequence_Sum

import "math"

type Deque []int

func (q Deque) Len() int {
	return len(q)
}

func (q Deque) IsEmpty() bool {
	return q.Len() == 0
}

func (q Deque) Front() int {
	return q[0]
}

func (q Deque) Back() int {
	return q[q.Len()-1]
}

func (q *Deque) Push(e int) {
	*q = append(*q, e)
}

func (q *Deque) PopBack() int {
	prev := *q
	back := prev.Back()
	*q = prev[0 : prev.Len()-1]
	return back
}

func (q *Deque) PopFront() int {
	prev := *q
	front := prev.Front()
	*q = prev[1:]
	return front
}

func constrainedSubsetSum(nums []int, k int) int {
	// 略有點像是找nums在windows size k的最大值
	// 但是這次nums[i]會被前面影響，so dp
	q := Deque{}
	n := len(nums)
	dp := make([]int, n)
	res := math.MinInt / 2
	for i := 0; i < n; i++ {
		if i > k && q.Front() == i-k-1 {
			q.PopFront()
		}
		dp[i] = 0
		// dp => 我要從前面開始 還是從自己開始往後加總，顯然 如果前面的加總是負數 我就從自己開始吧
		if !q.IsEmpty() {
			dp[i] = max(dp[q.Front()], 0)
		}
		dp[i] += nums[i]
		for !q.IsEmpty() && dp[i] >= dp[q.Back()] {
			q.PopBack()
		}
		q.Push(i)
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
