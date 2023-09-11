package _062_Longest_Repeating_Substring

import "testing"

func TestLongestRepeatingSubstring(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 0
		actual := longestRepeatingSubstring("abcd")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := 2
		actual := longestRepeatingSubstring("abbaba")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		expect := 3
		actual := longestRepeatingSubstring("aabcaabdaab")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
