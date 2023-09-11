package _156_Find_Substring_With_Given_Hash_Value

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	n := len(s)
	// Rolling hash 的基本題
	maxPower := int64(1)
	// maxPower => power ^ (k - 1)
	M := int64(modulo)
	for i := 0; i < k-1; i++ {
		maxPower = maxPower * int64(power) % M
	}
	ans := 0
	countingHash := int64(0)
	for i := n - 1; i >= 0; i-- {
		// Rolling
		if i+k < n {
			countingHash = countingHash - int64(s[i+k]-'a'+1)*maxPower%M
			countingHash = (countingHash + M) % M
		}
		countingHash = int64(power) * countingHash % M
		countingHash += int64(s[i] - 'a' + 1)
		countingHash = (countingHash + M) % M
		if countingHash == int64(hashValue) && i+k <= n {
			ans = i
		}
	}
	return s[ans : ans+k]

}
