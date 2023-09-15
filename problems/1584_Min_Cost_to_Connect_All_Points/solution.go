package _584_Min_Cost_to_Connect_All_Points

import "math"

type Edge struct {
	SrcIndex int
	DstIndex int
	Weight   int
}

type UnionFind struct {
	parents []int
	ranks   []int
}

func NewUnionFind(n int) UnionFind {
	parents := make([]int, n)
	ranks := make([]int, n)
	for i := range parents {
		parents[i] = i
		ranks[i] = i
	}
	return UnionFind{
		parents,
		ranks,
	}
}

func (u UnionFind) Find(x int) int {
	if u.parents[x] != x {
		u.parents[x] = u.Find(u.parents[x])
	}
	return u.parents[x]
}

func (u UnionFind) Union(x, y int) {
	rx := u.Find(x)
	ry := u.Find(y)
	if rx == ry {
		return
	}
	if rx < ry {
		u.parents[ry] = rx
		return
	}
	u.parents[rx] = ry
}

type PQ []Edge

func (q PQ) Len() int {
	return len(q)
}

func (q PQ) IsEmpty() bool {
	return q.Len() == 0
}

func (q PQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q PQ) Less(i, j int) bool {
	return q[i].Weight < q[j].Weight
}

func (q *PQ) Push(x interface{}) {
	*q = append(*q, x.(Edge))
}

func (q *PQ) Pop() interface{} {
	prev := *q
	top := prev[prev.Len()-1]
	prev = prev[0 : prev.Len()-1]
	*q = prev
	return top
}

// v3 Prime
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	if n == 1 {
		return 1
	}
	visited := make([]int, n)
	minDistance := make([]int, n)
	for i := range minDistance {
		minDistance[i] = math.MaxInt
	}
	minDistance[0] = 0
	ret := 0
	for i := 0; i < n; i++ {
		next := 0
		minD := math.MaxInt
		for j := 0; j < n; j++ {
			if visited[j] == 0 && minDistance[j] <= minD {
				minD = minDistance[j]
				next = j
			}
		}
		visited[next] = 1
		ret += minD
		for j := 0; j < n; j++ {
			if visited[j] == 0 {
				newNodeDistance := abs(points[j][0]-points[next][0]) + abs(points[j][1]-points[next][1])
				minDistance[j] = min(minDistance[j], newNodeDistance)
			}
		}
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Kruskal
//func minCostConnectPoints(points [][]int) int {
//v2 PQ
//n := len(points)
//uf := NewUnionFind(n)
//var calculateDistance func(i, j int) int
//edges := PQ{}
//heap.Init(&edges)
//count := 1
//totalWeight := 0
//calculateDistance = func(i, j int) int {
//	pi := points[i]
//	pj := points[j]
//	return abs(pi[0]-pj[0]) + abs(pi[1]-pj[1])
//}
//for i := 0; i < n; i++ {
//	for j := i + 1; j < n; j++ {
//		weight := calculateDistance(i, j)
//		heap.Push(&edges, Edge{
//			i,
//			j,
//			weight,
//		})
//	}
//}
//for !edges.IsEmpty() && count < n {
//	edge := heap.Pop(&edges).(Edge)
//	if uf.Find(edge.SrcIndex) == uf.Find(edge.DstIndex) {
//		continue
//	}
//	uf.Union(edge.SrcIndex, edge.DstIndex)
//	count++
//	totalWeight += edge.Weight
//}
//return totalWeight
// v1 sort
//sort.Slice(edges, func(i, j int) bool {
//	return edges[i].Weight < edges[j].Weight
//})
//i := 0
//
//for n > count {
//	edge := edges[i]
//	i++
//	if uf.Find(edge.SrcIndex) == uf.Find(edge.DstIndex) {
//		continue
//	}
//	count++
//	totalWeight += edge.Weight
//	uf.Union(edge.SrcIndex, edge.DstIndex)
//	if count == n {
//		break
//	}
//}
//return totalWeight

//}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
