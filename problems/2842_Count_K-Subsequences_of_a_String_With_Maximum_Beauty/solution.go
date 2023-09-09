package _842_Count_K_Subsequences_of_a_String_With_Maximum_Beauty

import "sort"

func countKSubsequencesWithMaxBeauty(s string, k int) int {
	const MOD = int64(1e9 + 7)
	occurrence := make(map[byte]int)
	n := len(s)
	sum := 0
	for i := 0; i < n; i++ {
		occurrence[s[i]]++
	}
	if len(occurrence) < k {
		return 0
	}
	var charOccurrence []int
	for _, occur := range occurrence {
		charOccurrence = append(charOccurrence, occur)
		sum += occur
	}
	sort.Slice(charOccurrence, func(i, j int) bool {
		return charOccurrence[i] > charOccurrence[j]
	})
	maxBeauty := 0
	total := int64(0)
	preSum := make([]int64, len(charOccurrence)+1)
	for i := 1; i < len(charOccurrence)+1; i++ {
		preSum[i] = preSum[i-1] + int64(charOccurrence[i-1])
	}
	for i := 0; i < k; i++ {
		maxBeauty += charOccurrence[i]
	}
	var dfs func(currPos, picked, currBeauty int, ret int64)
	dfs = func(currPos, picked, currBeauty int, ret int64) {
		if picked > k {
			return
		}
		if currBeauty == maxBeauty && picked == k {
			total = (total + ret) % MOD
		}
		if int64(currBeauty)+preSum[len(charOccurrence)]-preSum[currPos] < int64(maxBeauty) {
			return
		}
		for i := currPos; i < len(charOccurrence); i++ {
			dfs(i+1, picked+1, currBeauty+charOccurrence[i], (ret*int64(charOccurrence[i]))%MOD)
		}
	}
	dfs(0, 0, 0, int64(1))
	return int(total % MOD)
}
