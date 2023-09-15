package _370_Longest_Ideal_Subsequence

// 解2 dp => 以a,b,c,...,z為結尾的最長ideal subsequence O(n)
func longestIdealString(s string, k int) int {
	const A = int('a')
	n := len(s)
	var dp [26]int
	ret := 0
	for i := 0; i < n; i++ {
		c := int(s[i])
		t := 0
		lowerBound := max(c-A-k, 0)
		upperBound := min(c-A+k, 25)
		for j := lowerBound; j <= upperBound; j++ {
			t = max(t, dp[j]+1)
		}
		dp[c-A] = t
		ret = max(ret, t)
	}
	return ret
}

// 解1 DP為以i為結尾的最長ideal subsequence O(n^2)
//func longestIdealString(s string, k int) int {
//	const A = int('a')
//	n := len(s)
//	s = "#" + s
//	dp := make([]int, n+1)
//	for i := range dp {
//		dp[i] = 1
//	}
//	lastOccurrence := make(map[int]int)
//	ret := 0
//	for i := 1; i <= n; i++ {
//		c := int(s[i])
//		lowerBound := max(c-A-k, 0)
//		upperBound := min(c-A+k, 25)
//		for j := lowerBound; j <= upperBound; j++ {
//			if _, e := lastOccurrence[j]; e {
//				last := lastOccurrence[j]
//				dp[i] = max(dp[i], dp[last]+1)
//			}
//		}
//		lastOccurrence[c-A] = i
//		ret = max(ret, dp[i])
//	}
//	return ret
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
