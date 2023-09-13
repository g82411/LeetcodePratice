package _392_Longest_Happy_Prefix

import "testing"

func TestLongestPrefix(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := "l"
		actual := longestPrefix("level")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := "abab"
		actual := longestPrefix("ababab")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
