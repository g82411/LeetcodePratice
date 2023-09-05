package _483_Kth_Ancestor_of_a_Tree_Node

import "math"

type TreeAncestor struct {
	ancestors [][]int
}

func Constructor(n int, parent []int) TreeAncestor {
	maxStepOfAncestor := 20
	ancestors := make([][]int, n)
	for i := 0; i < n; i++ {
		d := make([]int, maxStepOfAncestor)
		for j := 0; j < maxStepOfAncestor; j++ {
			d[j] = -1
		}
		ancestors[i] = d
	}
	for i := 0; i < n; i++ {
		ancestors[i][0] = parent[i]
	}
	for j := 1; j < maxStepOfAncestor; j++ {
		for i := 0; i < n; i++ {
			node := ancestors[i][j-1]
			if node == -1 {
				ancestors[i][j] = ancestors[node][j-1]
			}
		}
	}
	return TreeAncestor{ancestors}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	maxStepToGetAncestor := int(math.Ceil(math.Log2(float64(k))))
	var bits []int
	for i := 0; i <= maxStepToGetAncestor; i++ {
		if ((k >> i) & 1) == 1 {
			bits = append(bits, i)
		}
	}
	cur := node
	for _, bit := range bits {
		cur = this.ancestors[cur][bit]
		if cur == -1 {
			return -1
		}
	}
	return cur
}
