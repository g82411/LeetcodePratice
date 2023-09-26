package _16_Remove_Duplicate_Letters

func removeDuplicateLetters(s string) string {
	// 注意這題不可以接將出現過的字母由小到大排序
	// 因為有些順序不可以更換
	// e.g bca => bca 因為沒有重複出現的字
	// 所謂最小是在這種情況 => baca => bac < bca

	remainOccur := make(map[byte]int)
	var res []byte
	for i := range s {
		remainOccur[s[i]]++
	}
	occur := make(map[byte]bool)
	for i := range s {
		if occur[s[i]] {
			remainOccur[s[i]]--
			continue
		}
		// 如果新的字母比較小 而且後面還出現過 可以先拔掉
		// e.g: bb a bbb
		for len(res) > 0 && res[len(res)-1] > s[i] && remainOccur[res[len(res)-1]] > 0 {
			occur[res[len(res)-1]] = false
			res = res[0 : len(res)-1]
		}
		res = append(res, s[i])
		remainOccur[s[i]]--
		occur[s[i]] = true
	}
	return string(res)
}
