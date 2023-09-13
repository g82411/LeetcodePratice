package _392_Longest_Happy_Prefix

func longestPrefix(s string) string {
	var build func(p string) []int
	build = func(p string) []int {
		n := len(p)
		pi := make([]int, n)
		j := 0
		for i := 1; i < n; i++ {
			for j > 0 && p[i] != p[j] {
				j = pi[j-1]
			}
			if p[i] == p[j] {
				j++
				pi[i] = j
			}
		}
		return pi
	}
	n := len(s)
	if n == 0 {
		return ""
	}
	pi := build(s)
	longest := pi[n-1]
	return s[0:longest]
}
