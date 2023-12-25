package _2977_Minimum_Cost_to_Convert_String_II

import "math"

type TrieNode struct {
	Next [26]*TrieNode
	Idx  int
}

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	uni := make(map[string]bool)
	for i := range original {
		uni[original[i]] = true
		uni[changed[i]] = true
	}

	// 對所有的點進行編碼 不然後面很難寫
	encodeMap := make(map[string]int)
	i := 0
	root := &TrieNode{}
	for k := range uni {
		encodeMap[k] = i
		// insert word to trie reversely
		node := root
		for j := 1; j <= len(k); j++ {
			charIdx := int(k[len(k)-j] - 'a')
			if node.Next[charIdx] == nil {
				node.Next[charIdx] = &TrieNode{
					Idx: -1,
				}
			}
			node = node.Next[charIdx]
		}
		node.Idx = i
		i++
	}
	n := len(encodeMap)
	d := make([][]int64, n)
	for i := range d {
		v := make([]int64, n)
		for j := range d {
			if i != j {
				v[j] = math.MaxInt64 / 2
			}
		}
		d[i] = v
	}
	// 開點
	for i := range original {
		d[encodeMap[original[i]]][encodeMap[changed[i]]] = min(
			d[encodeMap[original[i]]][encodeMap[changed[i]]],
			int64(cost[i]),
		)
	}
	// floyd法
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	m := len(source)
	source = "#" + source
	target = "#" + target
	dp := make([]int64, m+1)
	// dp[i] = min cost to convert source[:i] to target[:i]
	for i := 1; i <= m; i++ {
		dp[i] = math.MaxInt64 / 2
		// 如果source[i] == target[i]
		if source[i] == target[i] {
			dp[i] = dp[i-1]
		}
		sourceQuery := root
		targetQuery := root
		for j := i; j >= 1; j-- {
			nextSource := sourceQuery.Next[int(source[j]-'a')]
			nextTarget := targetQuery.Next[int(target[j]-'a')]
			if nextSource == nil {
				break
			}
			if nextTarget == nil {
				break
			}
			sourceQuery = nextSource
			targetQuery = nextTarget
			if sourceQuery.Idx == -1 || targetQuery.Idx == -1 {
				continue
			}
			dp[i] = min(dp[i], dp[j-1]+d[sourceQuery.Idx][targetQuery.Idx])
		}
	}
	if dp[m] == math.MaxInt64/2 {
		return -1
	}
	return dp[m]
}

// func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
//     uni := make(map[string]bool)
//     for i := range original {
//         uni[original[i]] = true
//         uni[changed[i]] = true
//     }

//     // 對所有的點進行編碼 不然後面很難寫
//     encodeMap := make(map[string]int)
//     i := 0
//     for k := range uni {
//         encodeMap[k] = i
//         i++
//     }
//     n := len(encodeMap)
//     d := make([][]int64, n)
//     for i := range d {
//         v := make([]int64, n)
//         for j := range d {
//             if i != j {
//                 v[j] = math.MaxInt64 / 2
//             }
//         }
//         d[i] = v
//     }
//     // 開點
//     for i := range original {
//         d[encodeMap[original[i]]][encodeMap[changed[i]]] = min(
//             d[encodeMap[original[i]]][encodeMap[changed[i]]],
//             int64(cost[i]),
//         )
//     }
//     // floyd法
//     for k := 0; k < n; k++ {
//         for i := 0; i < n; i++ {
//             for j := 0; j < n; j++ {
//                 d[i][j] = min(d[i][j], d[i][k] + d[k][j])
//             }
//         }
//     }
//     m := len(source)
//     source = "#" + source
//     target = "#" + target
//     dp := make([]int64, m + 1)
//     // dp[i] = min cost to convert source[:i] to target[:i]
//     for i := 1; i <= m; i++ {
//         dp[i] = math.MaxInt64 / 2
//         // 如果source[i] == target[i]
//         if source[i] == target[i] {
//             dp[i] = dp[i-1]
//         }
//         for j := i; j >= 1; j-- {
//             src := source[j:i+1]
//             dst := target[j:i+1]
//             idx1, eIdx1 := encodeMap[src]
//             idx2, eIdx2 := encodeMap[dst]
//             if ! (eIdx1 && eIdx2) {
//                 continue
//             }
//             dp[i] = min(dp[i], dp[j - 1] + d[idx1][idx2])
//         }
//     }
//     if dp[m] == math.MaxInt64 / 2 {
//         return -1
//     }
//     return dp[m]
// }
