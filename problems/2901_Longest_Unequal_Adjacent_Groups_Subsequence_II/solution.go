package _901_Longest_Unequal_Adjacent_Groups_Subsequence_II

func getHammingDistance(s1, s2 string) int {
	n := len(s1)
	diff := 0
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			continue
		}
		diff++
	}
	return diff
}

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// check is potential subseq
			if groups[i] == groups[j] {
				continue
			}
			if len(words[i]) != len(words[j]) {
				continue
			}
			if getHammingDistance(words[i], words[j]) > 1 {
				continue
			}
			// update longest sub seq
			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
	}
	maxLength := 0
	maxIndex := 0
	for i := 0; i < n; i++ {
		if dp[i] > maxLength {
			maxLength = dp[i]
			maxIndex = i
		}
	}
	var ans []string
	for maxIndex != -1 {
		ans = append([]string{words[maxIndex]}, ans...)
		maxIndex = prev[maxIndex]
	}
	return ans
}
