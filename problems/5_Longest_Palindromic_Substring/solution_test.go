package __Longest_Palindromic_Substring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expect := "bab"
		actual := longestPalindrome("babad")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		expect := "bb"
		actual := longestPalindrome("cbbd")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
