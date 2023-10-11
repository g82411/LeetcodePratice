package _682_Longest_Palindromic_Subsequence_II

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := longestPalindromeSubseq("bbabab")
		want := 4
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := longestPalindromeSubseq("dcbccacdb")
		want := 4
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
