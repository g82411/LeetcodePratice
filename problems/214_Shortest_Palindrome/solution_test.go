package _14_Shortest_Palindrome

import "testing"

func TestShortestPalindrome(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := "aaacecaaa"
		actual := shortestPalindrome("aacecaaa")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := "dcbabcd"
		actual := shortestPalindrome("abcd")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
