package _514_Path_with_Maximum_Probability

import "container/heap"

//
//type Edge struct {
//	next int
//	prob float64
//}
//
//type EdgeQueue []Edge
//
//func (q EdgeQueue) Len() int { return len(q) }
//
//func (q EdgeQueue) Less(i, j int) bool {
//	// 最大堆
//	return q[i].prob > q[j].prob
//}
//
//func (q EdgeQueue) Swap(i, j int) {
//	q[i], q[j] = q[j], q[i]
//}
//
//func (q *EdgeQueue) Push(x interface{}) {
//	edge := x.(Edge) // 進行類型斷言，將 interface{} 轉為 Edge
//	*q = append(*q, edge)
//}
//
//func (q *EdgeQueue) Pop() interface{} {
//	n := q.Len()
//	prev := *q
//	item := prev[n-1]
//	*q = prev[0 : n-1]
//	return item
//}
//
//func maxProbability(n int, edges [][]int, prob []float64, start int, end int) float64 {
//	adjust := make([][]Edge, n)
//	for i, edge := range edges {
//		p := prob[i]
//		s, e := edge[0], edge[1]
//		adjust[s] = append(adjust[s], Edge{next: e, prob: p})
//		adjust[e] = append(adjust[e], Edge{next: s, prob: p})
//	}
//
//	var maxHeap EdgeQueue
//	heap.Init(&maxHeap)
//	maxHeap.Push(Edge{next: start, prob: 1})
//	visited := make([]bool, n)
//
//	for maxHeap.Len() > 0 {
//		top := heap.Pop(&maxHeap).(Edge)
//		node, probability := top.next, top.prob
//
//		if visited[node] {
//			continue
//		}
//
//		if node == end {
//			return probability
//		}
//
//		visited[node] = true
//
//		for _, edge := range adjust[node] {
//			if !visited[edge.next] {
//				heap.Push(&maxHeap, Edge{next: edge.next, prob: edge.prob * probability})
//			}
//		}
//	}
//
//	return 0.0
//}

type Edge struct {
	Next int
	Prob float64
}

type EdgeQueue []Edge

func (q EdgeQueue) Len() int {
	return len(q)
}

func (q EdgeQueue) Less(i, j int) bool {
	// 最大堆
	return q[i].Prob > q[j].Prob
}

func (q EdgeQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *EdgeQueue) Push(edge interface{}) {
	*q = append(*q, edge.(Edge))
}

func (q *EdgeQueue) Pop() interface{} {
	n := q.Len()
	prev := *q
	item := prev[n-1]
	*q = prev[0 : n-1]
	return item
}
func maxProbability(n int, edges [][]int, prob []float64, start int, end int) float64 {
	adjust := make([][]Edge, n)
	// 注意怎麼套用一個heap 這比Dijkstra本身還難
	for i, edge := range edges {
		p := prob[i]
		s, e := edge[0], edge[1]
		adjust[s] = append(adjust[s], Edge{e, p})
		adjust[e] = append(adjust[e], Edge{s, p})
	}
	var maxHeap EdgeQueue
	heap.Init(&maxHeap)
	maxHeap.Push(Edge{start, 1})
	visited := make([]bool, n)

	// Dijkstra + BFS 的本體
	for maxHeap.Len() > 0 {
		// 原本的Dijkstra是每一次都找最短的那條路徑, 並期望總結果最小
		// 這邊會發現忽略了找最短的流程, 是因為在BFS的時候用的不是普通的Queue 而是PQ
		// 原本的PQ是找最小 可以稍微改一下變成maxHeap 就會變成題目要地找最大
		s := heap.Pop(&maxHeap).(Edge)
		if visited[s.Next] {
			continue
		}
		if s.Next == end {
			return s.Prob
		}
		visited[s.Next] = true
		for _, edge := range adjust[s.Next] {
			if visited[edge.Next] {
				continue
			}
			heap.Push(&maxHeap, Edge{edge.Next, edge.Prob * s.Prob})
		}
	}
	return 0.0
}
