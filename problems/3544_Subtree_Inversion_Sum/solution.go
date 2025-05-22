package _3544_Subtree_Inversion_Sum

import "math"

func subtreeInversionSum(edges [][]int, nums []int, k int) int64 {
	n := len(nums)
	adj := make([][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	memo := make([][][2]int64, n)
	for i := 0; i < n; i++ {
		memo[i] = make([][2]int64, k)
		for j := 0; j < k; j++ {
			memo[i][j] = [2]int64{math.MaxInt64 / 2, math.MaxInt64 / 2}
		}
	}
	var dfs func(cur, parent, cd, parity int) int64
	dfs = func(cur, parent, cd, parity int) int64 {
		cache := &memo[cur][cd][parity]
		if *cache != math.MaxInt64/2 {
			return *cache
		}
		res := int64(nums[cur] * (1 - parity*2))
		for _, next := range adj[cur] {
			if next == parent {
				continue
			}
			res += dfs(next, cur, max(cd-1, 0), parity)
		}
		if cd != 0 {
			*cache = res
			return res
		}
		inversion := int64(nums[cur] * (parity*2 - 1))
		for _, next := range adj[cur] {
			if next == parent {
				continue
			}
			inversion += dfs(next, cur, k-1, parity^1)
		}
		res = max(res, inversion)
		*cache = res
		return res
	}
	return dfs(0, -1, k-1, 0)
}
