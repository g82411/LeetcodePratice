package _858_Minimum_Edge_Reversals_So_Every_Node_Is_Reachable

// ReRoot
// 以0為root計算一個值
// 然後依序切換root 每個root理論上都只有0邊的+-1

type Edge struct {
	dst int
	dir int
}

func minEdgeReversals(n int, edges [][]int) []int {
	convertEdges := make([][]Edge, n)
	ret := make([]int, n)
	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		convertEdges[src] = append(convertEdges[src], Edge{
			dst,
			1,
		})
		convertEdges[dst] = append(convertEdges[dst], Edge{
			src,
			-1,
		})
	}
	var findDepthByPure func(cur, parent int) int
	var findDepthByCache func(cur, parent, reversal int)
	findDepthByPure = func(cur, parent int) int {
		ret := 0
		for _, edge := range convertEdges[cur] {
			if edge.dst == parent {
				continue
			}
			if edge.dir == 1 {
				ret += findDepthByPure(edge.dst, cur)
				continue
			}
			ret += findDepthByPure(edge.dst, cur) + 1
		}
		return ret
	}
	findDepthByCache = func(cur, parent, reversal int) {
		ret[cur] = reversal
		for _, edge := range convertEdges[cur] {
			if edge.dst == parent {
				continue
			}
			// 這邊有個小問題, 為什麼順邊要 + 1
			// 我的理解是這樣 因為這個dfs是為了去計算以next 為root 總共要反轉幾個邊
			// 則對n 來說是順邊的話 n + 1就會是反邊
			findDepthByCache(edge.dst, cur, reversal+edge.dir)
		}
	}
	initialNum := findDepthByPure(0, -1)
	findDepthByCache(0, -1, initialNum)

	return ret
}
