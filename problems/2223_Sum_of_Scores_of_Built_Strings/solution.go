package _223_Sum_of_Scores_of_Built_Strings

// 注意這題要小心 abcXXabcXXXXX 時會找到abcXXXXX 這時候會有一個prefix length 3
// 如果是從前往後取 跟 從後往前取比較 會少算到這種case
// 先做個binary saerch 跟他爆了
func sumScores(s string) int64 {
	n := len(s)
	var getSubStringHash func(from, to int) uint64
	prefixHash := make([]uint64, n)
	power := make([]uint64, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			// TODO: 確認一下這邊為什麼不用 + 1
			prefixHash[i] = uint64(s[i] - 'a')
			power[i] = 1
			continue
		}
		prefixHash[i] = prefixHash[i-1]*26 + uint64(s[i]-'a')
		power[i] = power[i-1] * 26
	}
	getSubStringHash = func(from, to int) uint64 {
		ret := prefixHash[from+to-1]
		if from > 0 {
			ret -= prefixHash[from-1] * power[to]
		}
		return ret
	}
	ret := int64(0)
	for i := n - 1; i >= 0; i-- {
		if s[i] != s[0] {
			// 顯然不是prefix 快滾
			continue
		}
		l := 1
		r := n - i
		for l <= r {
			mid := l + (r-l)/2
			if getSubStringHash(i, mid) != prefixHash[mid-1] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		ret += int64(r)
	}
	return ret
}
