package _846_Minimum_Edge_Weight_Equilibrium_Queries_in_a_Tree

import (
	"math"
)

type Edge struct {
	Next   int
	Weight int
}

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	// 這題很難 建議多看幾次 觀念很多
	// 1. 一棵樹 意味著可以拿任意節點做為根
	// 2. 題目問的是最小變更的次數 => 總共有多少條邊 - weight 頻次最高的邊
	// 在已知共同祖先的情況下 p -> q 的邊 = root -> p + root -> q - root -> pq的祖先
	// 10005式題目給的點數
	// 從i開始到有哪些邊
	treeEdge := make([][]Edge, 10005)
	// 從root -> i, weight = w的邊共有幾條
	weightFrequency := make([][27]int, 10005)
	ancestors := make([][18]int, 10005)
	parents := make([]int, 10005)
	levels := make([]int, 10005)
	buildTreeEdgeTable := func() {
		for _, edge := range edges {
			s, d, w := edge[0], edge[1], edge[2]
			treeEdge[s] = append(treeEdge[s], Edge{d, w})
			treeEdge[d] = append(treeEdge[d], Edge{s, w})
		}
	}
	// binary lifting
	var buildAncestorsTable func()
	var weightCount [27]int
	var treeTraversal func(current, parent, level int)
	var getLCA func(node1, node2 int) int
	var getKthAncestor func(node, k int) int
	var getMinimumChangeBetweenNode func(node1, node2 int) int
	treeTraversal = func(current, parent, level int) {
		parents[current] = parent
		levels[current] = level
		for _, edge := range treeEdge[current] {
			if edge.Next == parent {
				continue
			}
			// Counting child weight frequency
			weightCount[edge.Weight]++
			for i := 1; i <= 26; i++ {
				weightFrequency[edge.Next][i] = weightCount[i]
			}
			treeTraversal(edge.Next, current, level+1)
			// 夭壽喔這邊還有回朔法 殺了我
			weightCount[edge.Weight]--
		}
	}
	//getLCA1 := func(node1, node2 int) int {
	//	n1, n2 := node1, node2
	//	for n1 != n2 {
	//		if levels[n1] > levels[n2] {
	//			n1 = parents[n1]
	//		} else if levels[n1] < levels[n2] {
	//			n2 = parents[n2]
	//		} else if n1 != n2 {
	//			n1 = parents[n1]
	//			n2 = parents[n2]
	//		}
	//	}
	//	return n1
	//}
	//binary lifting
	buildAncestorsTable = func() {
		maxLiftStep := int(math.Ceil(math.Log2(float64(n))))
		for i := 0; i < n; i++ {
			ancestors[i][0] = parents[i]
		}
		// avoid null pointer
		ancestors[0][0] = 0
		for j := 1; j <= maxLiftStep; j++ {
			for i := 0; i < n; i++ {
				ancestors[i][j] = ancestors[ancestors[i][j-1]][j-1]
			}
		}
	}
	getKthAncestor = func(node, k int) int {
		cur := node
		for j := 0; j <= 17; j++ {
			if ((k >> j) & 1) == 1 {
				cur = ancestors[cur][j]
			}
		}
		return cur
	}
	getLCA = func(node1, node2 int) int {
		n1, n2 := node1, node2
		for levels[n1] != levels[n2] {
			if levels[n1] > levels[n2] {
				n1 = getKthAncestor(n1, levels[n1]-levels[n2])
			} else if levels[n1] < levels[n2] {
				n2 = getKthAncestor(n2, levels[n2]-levels[n1])
			}
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
	getMinimumChangeBetweenNode = func(node1, node2 int) int {
		sumOfEdge := 0
		maxEdgeFrequency := 0
		lca := getLCA(node1, node2)
		for i := 1; i <= 26; i++ {
			edgeCount := weightFrequency[node1][i] + weightFrequency[node2][i] - 2*weightFrequency[lca][i]
			sumOfEdge += edgeCount
			maxEdgeFrequency = max(maxEdgeFrequency, edgeCount)
		}
		return sumOfEdge - maxEdgeFrequency
	}
	ret := []int{}
	buildTreeEdgeTable()
	treeTraversal(0, -1, 0)
	buildAncestorsTable()
	for _, query := range queries {
		minimumChange := getMinimumChangeBetweenNode(query[0], query[1])
		ret = append(ret, minimumChange)
	}
	return ret

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
