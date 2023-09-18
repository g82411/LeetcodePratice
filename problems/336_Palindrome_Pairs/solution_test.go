package _36_Palindrome_Pairs

import "testing"

func TestPalindromePairs(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expect := [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}
		actual := palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"})
		if len(expect) != len(actual) {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
