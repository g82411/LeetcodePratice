package _2920_Maximum_Points_After_Collecting_Coins_From_All_Nodes

import "math"

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	memos := make([][]int, n)
	neighbors := make([][]int, n)
	for i := range memos {
		memos[i] = make([]int, 15)
		neighbors[i] = make([]int, 0)
		for j := range memos[i] {
			memos[i][j] = math.MinInt
		}
	}
	for _, e := range edges {
		src, dst := e[0], e[1]
		neighbors[src] = append(neighbors[src], dst)
		neighbors[dst] = append(neighbors[dst], src)
	}

	var dfs func(i, shift, parent int) int
	dfs = func(i, shift, parent int) int {
		shift = min(shift, 14)
		if memos[i][shift] > math.MinInt {
			return memos[i][shift]
		}
		s1 := (coins[i] >> shift) - k
		s2 := (coins[i] >> shift) / 2
		for _, dst := range neighbors[i] {
			if dst == parent {
				continue
			}
			s1 += dfs(dst, shift, i)
			s2 += dfs(dst, shift+1, i)
		}
		memos[i][shift] = max(s1, s2)
		return memos[i][shift]
	}
	return dfs(0, 0, -1)

}
