package _87_Scramble_String

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s1) {
		return false
	}
	// 先比較詞頻，如果詞頻不一致，就false
	var s1Letters, s2Letters [26]int
	for _, l := range s1 {
		s1Letters[int(l-'a')]++
	}
	for _, l := range s2 {
		s2Letters[int(l-'a')]++
	}
	if s1Letters != s2Letters {
		return false
	}
	// dfs
	memo := make(map[string]bool)
	var test func(s, t string) bool
	test = func(s, t string) bool {
		// 終止條件，s == t || lens(s) == 1 => 照抄題目
		if s == t || len(s) == 1 {
			return true
		}
		n := len(s)
		// cache
		key := s + t
		if val, ok := memo[key]; ok {
			return val
		}

		var sl1, sl2, tl [26]int
		for i := 1; i < n; i++ {
			// 找分割點
			sl1[s[i-1]-'a']++
			sl2[s[n-i]-'a']++
			tl[t[i-1]-'a']++
			// 如果前半段詞頻跟t的一致，那就有機會是一個可行轉換，去比較前半段的轉換是否成立，後半段的轉換是否成立
			if sl1 == tl && test(s[:i], t[:i]) && test(s[i:], t[i:]) {
				memo[key] = true
				return true
			}
			// 如果後半段詞頻跟t的一致，那就有機會是一個可行轉換，去比較後半段的轉換是否成立，後半段的轉換是否成立
			if sl2 == tl && test(s[n-i:], t[:i]) && test(s[:n-i], t[i:]) {
				memo[key] = true
				return true
			}
			memo[key] = false
		}

		return false
	}

	return test(s1, s2)
}
