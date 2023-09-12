package _31_Palindrome_Partitioning

import "testing"

func TestPartition(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		actual := partition("aab")
		expect := [][]string{{"a", "a", "b"}, {"aa", "b"}}
		if len(actual) != len(expect) {
			t.Errorf("expect %v actual %v", expect, actual)
		}
	})
}
