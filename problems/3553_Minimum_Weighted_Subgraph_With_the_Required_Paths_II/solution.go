package _3553_Minimum_Weighted_Subgraph_With_the_Required_Paths_II

import (
	"math"
)

type Edge struct {
	Next   int
	Weight int
}

func minimumWeight(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	nexts := make([][]Edge, n)
	for _, edge := range edges {
		s, d, w := edge[0], edge[1], edge[2]
		nexts[s] = append(nexts[s], Edge{d, w})
		nexts[d] = append(nexts[d], Edge{s, w})
	}
	ancestors := make([][18]int, n)
	// 17 -> 2^17 = 131072
	// max edges = 100000 so max nodes => 100000 + 1
	parents := make([]int, n)
	levels := make([]int, n)
	var treeTraversal func(current, parent, level int)
	treeTraversal = func(current, parent, level int) {
		parents[current] = parent
		levels[current] = level
		for _, edge := range nexts[current] {
			if edge.Next == parent {
				continue
			}
			treeTraversal(edge.Next, current, level+1)
		}
	}
	treeTraversal(0, -1, 0)
	distance := make([]int, n)
	var findAllDistanceFromRoot func()
	var buildAncestorsTable func()
	var getLCA func(node1, node2 int) int
	var getKthAncestor func(node, k int) int
	// 再講一次binary lifting的原理
	// 找第k個祖先 k = 2^0 + 2^1 + 2^2 + ... + 2^n
	// 所以可以用二進位的方式來表示
	// 例如 13 = 1101 => 2^0 + 2^2 + 2^3
	// 這個func先來找找某個digit的祖先
	buildAncestorsTable = func() {
		maxLiftStep := int(math.Ceil(math.Log2(float64(n))))
		for i := 0; i < n; i++ {
			ancestors[i][0] = parents[i]
		}
		ancestors[0][0] = 0
		for j := 1; j <= maxLiftStep; j++ {
			for i := 0; i < n; i++ {
				ancestors[i][j] = ancestors[ancestors[i][j-1]][j-1]
			}
		}
	}
	buildAncestorsTable()
	getKthAncestor = func(node, k int) int {
		cur := node
		for i := 0; i < 18; i++ {
			if ((k >> i) & 1) == 1 {
				cur = ancestors[cur][i]
			}
		}
		return cur
	}
	getLCA = func(node1, node2 int) int {
		n1, n2 := node1, node2
		if levels[n1] < levels[n2] {
			n1, n2 = n2, n1
		}
		for levels[n1] != levels[n2] {
			if levels[n1] > levels[n2] {
				n1 = getKthAncestor(n1, levels[n1]-levels[n2])
				continue
			}
			n2 = getKthAncestor(n2, levels[n2]-levels[n1])
		}
		l, r := 0, levels[n1]
		for l < r {
			m := l + (r-l)/2
			a1 := getKthAncestor(n1, m)
			a2 := getKthAncestor(n2, m)
			if a1 == a2 {
				r = m
			} else {
				l = m + 1
			}
		}
		return getKthAncestor(n1, l)
	}

	findAllDistanceFromRoot = func() {
		var dfs func(current, parent int)
		dfs = func(current, parent int) {
			for _, edge := range nexts[current] {
				if edge.Next == parent {
					continue
				}
				distance[edge.Next] = distance[current] + edge.Weight
				dfs(edge.Next, current)
			}
		}
		dfs(0, -1)
	}

	findAllDistanceFromRoot()

	getDis := func(node1, node2 int) int {
		lca := getLCA(node1, node2)
		return distance[node1] + distance[node2] - 2*distance[lca]
	}

	ret := make([]int, len(queries))
	for i, query := range queries {
		node1, node2, node3 := query[0], query[1], query[2]
		ret[i] = getDis(node1, node2) - getDis(node1, node3)
	}
	return ret
}
