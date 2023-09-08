package _599_Make_the_Prefix_Sum_Non_negative

import "container/heap"

type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(e interface{}) {
	*m = append(*m, e.(int))
}

func (m *MinHeap) Pop() interface{} {
	// 注意他叫Push & Pop, 所以要彈出最後一個元素 實作跟stack 一樣...
	prev := *m
	n := len(prev)
	element := prev[n-1]
	*m = prev[0 : n-1]
	return element
}

func makePrefSumNonNegative(nums []int) int {
	var minHeap MinHeap
	heap.Init(&minHeap)
	prefixSum := 0
	steps := 0
	for _, num := range nums {
		prefixSum += num
		if num < 0 {
			heap.Push(&minHeap, abs(num))
		}
		if prefixSum < 0 {
			m := heap.Pop(&minHeap)
			prefixSum += m.(int)
			steps++
		}
	}
	return steps
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
