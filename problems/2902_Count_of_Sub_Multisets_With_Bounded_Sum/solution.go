package _902_Count_of_Sub_Multisets_With_Bounded_Sum

// 看到想到背包問題,
// dp[i][j] = dp[i-1][j] + dp[i][j-v]
// 但是注意: A sub-multiset is an unordered collection of elements of the array in which a given value x can occur 0, 1, ..., occ[x] times, where occ[x] is the number of occurrences of x in the array.
// 故dp[i][j] = 用i個element 填滿j的方法 很像背包但不是
//
//	for i := 1; i <= M; i++ {
//	    for j := 0; j <= C; j++ {
//	        dp[i][j] = dp[i-1][j] + dp[i-1][j-vi] + dp[i-1][j-vi * 2] + ... + dp[i - 1][j-ai*vi]
//	    }
//	}
//
// yoo因為dp[i][j] = dp[i-1][j] + dp[i-1][j-vi] + dp[i-1][j-vi * 2] + ... + dp[i - 1][j-ai*vi]
// 醬子要ai次
// yoo dp[i][j] =  dp[i-1][j-vi] + dp[i-1][j-vi*2] + dp[i-1][j-vi * 3] + ... + dp[i - 1][j-(ai+1)*vi]
// dp[i][j] - dp[i][j-vi] = dp[i-1][j] + dp[i - 1][j-(ai+1)*vi]
// dp[i][j] = dp[i][j-vi] + dp[i-1][j] + dp[i - 1][j-(ai+1)*vi]
func countSubMultisets(nums []int, l int, r int) int {
	M := int(1e9) + 7
	zeroOccur := 0
	occur := make(map[int]int)
	for _, num := range nums {
		if num == 0 {
			zeroOccur++
			continue
		}
		occur[num]++
	}
	freq := [][]int{{0, 0}}
	for k, v := range occur {
		freq = append(freq, []int{k, v})
	}
	helper := func(C int) int {
		if C < 0 {
			return 0
		}
		m := len(freq) - 1
		dp := make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, C+1)
		}
		dp[0][0] = 1
		for i := 1; i <= m; i++ {
			value, amount := freq[i][0], freq[i][1]
			for j := 0; j <= C; j++ {
				dp[i][j] = dp[i-1][j]
				if j >= value {
					dp[i][j] += dp[i][j-value]
				}
				if j >= value*(amount+1) {
					dp[i][j] -= dp[i-1][j-value*(amount+1)]
				}
				dp[i][j] = (dp[i][j] + M) % M
			}
		}
		ret := 0
		for j := 0; j <= C; j++ {
			ret = (ret + dp[m][j]) % M
		}
		return (max(1, ret) * (zeroOccur + 1)) % M
	}
	// fmt.Println(helper(r) , helper(l - 1))
	return (helper(r) - helper(l-1) + M) % M
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
