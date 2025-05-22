package _3362_Zero_Array_Transformation_III

import (
	"container/heap"
	"sort"
)

type PQ []int

func (q PQ) Len() int {
	return len(q)
}

func (q PQ) IsEmpty() bool {
	return len(q) == 0
}

func (q PQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q PQ) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q *PQ) Push(edge interface{}) {
	*q = append(*q, edge.(int))
}

func (q *PQ) Pop() interface{} {
	prev := *q
	top := prev[q.Len()-1]
	*q = prev[0 : q.Len()-1]
	return top
}

func maxRemoval(nums []int, queries [][]int) int {
	// sort
	sort.Slice(queries, func(i, j int) bool {
		if queries[i][0] == queries[j][0] {
			return queries[i][1] > queries[j][1]
		}
		return queries[i][0] < queries[j][0]
	})
	n := len(nums)
	pq := &PQ{}
	heap.Init(pq)
	delta := make([]int, n+1)
	operations := 0
	for i, j := 0, 0; i < n; i++ {
		operations += delta[i]
		for j < len(queries) && queries[j][0] <= i {
			heap.Push(pq, queries[j][1])
			j++
		}
		for operations < nums[i] && pq.Len() > 0 && (*pq)[0] >= i {
			operations++
			delta[heap.Pop(pq).(int)+1]--
		}
		if operations >= nums[i] {
			return -1
		}
	}
	return pq.Len()
}
