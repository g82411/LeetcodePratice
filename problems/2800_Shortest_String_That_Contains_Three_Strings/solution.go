package _800_Shortest_String_That_Contains_Three_Strings

import "strings"

// 484有一種可能
// a = xxxx
// b = xxxy
// a + b xxxxy就可以達到最小
// 意思是 找到max{suffix(a) = prefix(b)}的組合 就口以組合出最長的merge

func minimumString(a string, b string, c string) string {
	mergeString := func(a, b string) string {
		if strings.Contains(b, a) {
			return b
		}
		for overlapSize := 0; overlapSize < len(a); overlapSize++ {
			aSuffix := a[overlapSize:]
			bPrefix := b[:min(len(aSuffix), len(b))]
			if aSuffix == bPrefix {
				bSuffix := b[len(bPrefix):]
				return a + bSuffix
			}
		}
		return a + b
	}
	enums := [][3]string{
		[3]string{a, b, c},
		[3]string{a, c, b},
		[3]string{b, a, c},
		[3]string{b, c, a},
		[3]string{c, b, a},
		[3]string{c, a, b},
	}
	ans := ""
	for _, e := range enums {
		x, y, z := e[0], e[1], e[2]
		candidate := mergeString(x, mergeString(y, z))
		if ans == "" {
			ans = candidate
			continue
		}
		if len(ans) > len(candidate) {
			ans = candidate
			continue
		}
		if len(ans) == len(candidate) && ans > candidate {
			ans = candidate
			continue
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
