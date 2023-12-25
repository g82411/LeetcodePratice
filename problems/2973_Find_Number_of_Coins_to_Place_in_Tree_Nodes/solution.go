package _2973_Find_Number_of_Coins_to_Place_in_Tree_Nodes

import "sort"

func placedCoins(edges [][]int, cost []int) []int64 {
	var next [20005][]int
	var children [20005][]int64
	rets := make([]int64, len(cost))
	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		next[src] = append(next[src], dst)
		next[dst] = append(next[dst], src)
	}
	var dfs func(cur, parent int)
	dfs = func(cur, parent int) {
		var tmp []int64
		for _, nxt := range next[cur] {
			if nxt == parent {
				continue
			}
			dfs(nxt, cur)
			// 把子節點的cost一個一個加進去temp
			for _, x := range children[nxt] {
				tmp = append(tmp, int64(x))
			}
		}
		tmp = append(tmp, int64(cost[cur]))
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		n := len(tmp)
		if n < 3 {
			rets[cur] = 1
		} else {
			// 最大值有兩種口能，一種是很基本的最大值3個相乘，另外一個是最小的兩個負數相乘 再乘上最大的正數
			rets[cur] = max(
				0,
				tmp[0]*tmp[1]*tmp[n-1],
				tmp[n-1]*tmp[n-2]*tmp[n-3],
			)
		}
		// 這邊注意一種coin只能用一次
		if n >= 1 {
			children[cur] = append(children[cur], tmp[0])
		}
		if n >= 2 {
			children[cur] = append(children[cur], tmp[1])
		}
		if n >= 5 {
			children[cur] = append(children[cur], tmp[n-3])
		}
		if n >= 4 {
			children[cur] = append(children[cur], tmp[n-2])
		}
		if n >= 3 {
			children[cur] = append(children[cur], tmp[n-1])
		}
	}
	dfs(0, -1)
	return rets
}
