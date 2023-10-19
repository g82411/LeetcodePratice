package _904_Shortest_and_Lexicographically_Smallest_Beautiful_String

import (
	"math"
	"sort"
)

// 就恩...挺簡單的
func shortestBeautifulSubstring(s string, k int) string {
	l := 0
	n := len(s)
	one := 0
	var ans []string
	shortest := math.MaxInt / 2
	for r := 0; r < n; r++ {
		if s[r] == '1' {
			one++
		}
		for one >= k && l <= r {

			if one == k {
				length := r - l + 1
				if length < shortest {
					shortest = length
					ans = ans[:0]
				}
				if length == shortest {
					ans = append(ans, s[l:r+1])
				}
			}
			if s[l] == '1' {
				one--
			}
			l++
		}
	}
	sort.Strings(ans)
	if len(ans) == 0 {
		return ""
	}
	return ans[0]
}
