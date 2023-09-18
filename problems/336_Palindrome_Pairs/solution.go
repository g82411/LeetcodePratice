package _36_Palindrome_Pairs

// 回文組合法
// 如果s[i:j]是回文 則要加前後墜 不可能符合題宜
// 所以s[i:]是回文, 則找一個reverse 過的word = s[:i] word+ s 就可以是回文
// 另外一種可能是s[:-i]是回文 則找一個reverse 過的word = s[-i:] s + word就是回文
func palindromePairs(words []string) [][]int {
	candidate := make(map[string]int)
	for i, word := range words {
		key := reverse(word)
		candidate[key] = i
	}
	result := make(map[[2]int]bool)
	for i, word := range words {
		n := len(word)
		for j := 0; j <= n; j++ {
			prefix := word[:j]
			suffix := word[n-j:]
			if idx, e := candidate[prefix]; e && idx != i && isPalindrome(word[j:]) {
				result[[2]int{i, idx}] = true
			}
			if idx, e := candidate[suffix]; e && idx != i && isPalindrome(word[:n-j]) {
				result[[2]int{idx, i}] = true
			}
		}
	}
	var ans [][]int
	for pair := range result {
		n := pair
		ans = append(ans, n[:])
	}
	return ans
}

func isPalindrome(word string) bool {
	for l, r := 0, len(word)-1; l < r; l, r = l+1, r-1 {
		if word[l] != word[r] {
			return false
		}
	}

	return true
}

func reverse(word string) string {
	str := []byte(word)

	for l, r := 0, len(word)-1; l < r; l, r = l+1, r-1 {
		str[l], str[r] = str[r], str[l]
	}

	return string(str)
}
