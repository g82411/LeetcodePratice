package _554_Strings_Differ_by_One_Character

// 還是一個小問題 hashTable 他不香嗎?
// 詳細https://leetcode.com/problems/strings-differ-by-one-character/solutions/1794632/golang-map/
//
//	直接爆ram 所以rolling hash還是有其好處的
func differByOne(dict []string) bool {
	base := uint64(27)
	var rollingHash func(s string) uint64
	rollingHash = func(s string) uint64 {
		hash := uint64(0)
		for i := 0; i < len(s); i++ {
			hash = hash*base + uint64(s[i]-'a'+1)
		}
		return hash
	}
	hashes := []uint64{}
	for _, s := range dict {
		hashes = append(hashes, rollingHash(s))
	}
	n := len(dict)
	stringLen := len(dict[0])
	p := uint64(1)
	for j := stringLen - 1; j >= 0; j-- {
		set := make(map[uint64]int)
		for i := 0; i < n; i++ {
			new_hash := hashes[i] - uint64(dict[i][j]-'a'+1)*p
			if _, e := set[new_hash]; e {
				return true
			}
			set[new_hash]++
		}
		p *= base
	}
	return false
}
