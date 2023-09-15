package _370_Longest_Ideal_Subsequence

import "testing"

func TestLongestIdealString(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 4
		actual := longestIdealString("acfgbd", 2)
		if expect != actual {
			t.Errorf("Expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := 4
		actual := longestIdealString("abcd", 3)
		if expect != actual {
			t.Errorf("Expect: %v, actual: %v", expect, actual)
		}
	})
}
