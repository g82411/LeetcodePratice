package _657_Determine_if_Two_Strings_Are_Close

import "sort"

func closeStrings(word1 string, word2 string) bool {
	frequency1 := make([]int, 26)
	frequency2 := make([]int, 26)
	charSet1 := 0
	charSet2 := 0
	if len(word1) != len(word2) {
		return false
	}
	for i, c := range word1 {
		c1 := int(c - 'a')
		c2 := int(rune(word2[i]) - 'a')
		charSet1 = charSet1 | 1<<c1
		charSet2 = charSet2 | 1<<c2
		frequency1[c1]++
		frequency2[c2]++
	}
	sort.Ints(frequency1)
	sort.Ints(frequency2)
	for i := 0; i < 26; i++ {
		if frequency1[i] != frequency2[i] {
			return false
		}
	}
	return charSet1 == charSet2
}
