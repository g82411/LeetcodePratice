package _43_Network_Delay_Time

import "container/heap"

type Edge struct {
	Point  int
	Weight int
}

type PQ []Edge

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq *PQ) Push(e interface{}) {
	*pq = append(*pq, e.(Edge))
}

func (pq *PQ) Pop() interface{} {
	prev := *pq
	top := prev[prev.Len()-1]
	*pq = prev[0 : prev.Len()-1]
	return top
}

func networkDelayTime(times [][]int, n int, k int) int {
	visited := make(map[int]bool)
	edges := make(map[int][]Edge)
	res := 0
	for _, time := range times {
		curr, next, weight := time[0], time[1], time[2]
		edges[curr] = append(edges[curr], Edge{
			next,
			weight,
		})
	}
	pq := PQ{}
	heap.Init(&pq)
	heap.Push(&pq, Edge{
		k,
		0,
	})
	for !pq.IsEmpty() {
		curr := heap.Pop(&pq).(Edge)
		currPoint, currWeight := curr.Point, curr.Weight
		if visited[currPoint] {
			continue
		}
		visited[currPoint] = true
		res = max(res, currWeight)
		for _, edge := range edges[currPoint] {
			if visited[edge.Point] {
				continue
			}
			heap.Push(&pq, Edge{
				edge.Point,
				edge.Weight + currWeight,
			})
		}
	}
	if len(visited) == n {
		return res
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
