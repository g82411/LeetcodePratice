package _489_Find_Critical_and_Pseudo_Critical_Edges_in_Minimum_Spanning_Tree

import (
	"math"
	"sort"
)

type Edge struct {
	Src    int
	Dst    int
	Weight int
	OriIdx int
}

type UnionFind struct {
	parents []int
}

func NewUnionFind(n int) UnionFind {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	return UnionFind{
		parents,
	}
}

func (u UnionFind) Find(x int) int {
	if x != u.parents[x] {
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
	if rx > ry {
		u.parents[rx] = ry
		return
	}
	u.parents[ry] = rx
}

func kruskalFindMSTWeight(n, ignoreIdx int, edges []Edge) int {
	w := 0
	uf := NewUnionFind(n)
	for _, edge := range edges {
		if uf.Find(edge.Src) == uf.Find(edge.Dst) {
			continue
		}
		if edge.OriIdx == ignoreIdx {
			continue
		}
		uf.Union(edge.Src, edge.Dst)
		w += edge.Weight
	}
	for i := 0; i < n; i++ {
		if uf.Find(i) != uf.Find(0) {
			return math.MaxInt
		}
	}
	return w
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	// There are two set, one is the edge which we remove it, will make MST larger, or even can't still be a MST
	// And other is such edge we can remove some of edge in set, and build MST with the same weight
	// So we can force to use some edge and check if weight is still the same
	largeSet := make(map[int]bool)
	sameSet := make(map[int]bool)
	var newEdges []Edge
	for i, edge := range edges {
		newEdges = append(newEdges, Edge{
			edge[0],
			edge[1],
			edge[2],
			i,
		})
	}
	sort.Slice(newEdges, func(i, j int) bool {
		return newEdges[i].Weight < newEdges[j].Weight
	})
	minWeight := kruskalFindMSTWeight(n, -1, newEdges)
	for _, edge := range newEdges {
		m := kruskalFindMSTWeight(n, edge.OriIdx, newEdges)
		if m > minWeight {
			largeSet[edge.OriIdx] = true
		}
	}
	for _, edge := range newEdges {
		m := kruskalFindMSTWeight(n, edge.OriIdx, newEdges)
		if m > minWeight {
			largeSet[edge.OriIdx] = true
		}
	}

	for _, edge := range newEdges {
		if _, e := largeSet[edge.OriIdx]; e {
			continue
		}
		newEdges = append([]Edge{edge}, newEdges...)
		m := kruskalFindMSTWeight(n, -1, newEdges)
		if m == minWeight {
			sameSet[edge.OriIdx] = true
		}
		newEdges = newEdges[1:]

	}
	var res [][]int
	var same []int
	var large []int
	for k := range largeSet {
		large = append(large, k)
	}
	for k := range sameSet {
		same = append(same, k)
	}
	res = append(res, large)
	res = append(res, same)
	return res

}
