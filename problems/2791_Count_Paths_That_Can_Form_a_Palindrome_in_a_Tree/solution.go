package _791_Count_Paths_That_Can_Form_a_Palindrome_in_a_Tree

// 處理root -> u == root -> v
// 找到LCA a 使au = av
// 則找到u v 字符出現頻率一致
// 另外 我們不關心u v的順序 只關心u v的頻率

type Node struct {
	Alpha  byte
	Number int
}

func countPalindromePaths(parent []int, s string) int64 {
	n := len(parent)
	edges := make(map[int][]Node)
	var dfs func(curr, parent, state int)
	ret := int64(0)
	count := make(map[int]int64)
	for i := 1; i < n; i++ {
		edges[parent[i]] = append(edges[parent[i]], Node{
			s[i],
			i,
		})
	}
	dfs = func(curr, parent, state int) {
		// 如果這個state 存在過，那這次就找到第二條同樣的state
		if _, e := count[state]; e {
			ret += count[state]
		}
		// 如果這個state 跟curr state 差一個count, 那也可以加進去變成回文好棒
		for i := 0; i < 26; i++ {
			s := state ^ (1 << i)
			if _, e := count[s]; e {
				ret += count[s]
			}
		}
		count[state]++
		for _, next := range edges[curr] {
			if next.Number == parent {
				continue
			}
			dfs(next.Number, curr, state^(1<<(next.Alpha-'a')))
		}
	}
	dfs(0, -1, 0)
	return ret
}
