package _31_Palindrome_Partitioning

// 切一刀 看左邊 看右邊
func partition(s string) [][]string {
	var result [][]string
	var candidate []string
	var isPalindrome func(start, end int) bool
	var dfs func(start int)
	isPalindrome = func(start, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}
	dfs = func(start int) {
		if start == len(s) {
			c := make([]string, len(candidate))
			copy(c, candidate)
			result = append(result, c)
			return
		}
		for end := start; end < len(s); end++ {
			if !isPalindrome(start, end) {
				continue
			}
			palindrome := s[start : end+1]
			candidate = append(candidate, palindrome)
			dfs(end + 1)
			candidate = candidate[0 : len(candidate)-1]
		}
	}
	dfs(0)
	return result
}
