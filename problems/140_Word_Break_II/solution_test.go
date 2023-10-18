package _40_Word_Break_II

import "testing"

func TestWordBreak(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := []string{"cat sand dog", "cats and dog"}
		actual := wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
		if !equal(expect, actual) {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := []string{}
		actual := wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
		if !equal(expect, actual) {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		expect := []string{}
		actual := wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
		if !equal(expect, actual) {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 4", func(t *testing.T) {
		expect := []string{"aaaa aaa", "aaa aaaa"}
		actual := wordBreak("aaaaaaa", []string{"aaaa", "aaa"})
		if !equal(expect, actual) {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}

func equal(expect []string, actual []string) bool {
	if len(expect) != len(actual) {
		return false
	}
	for i := range expect {
		if expect[i] != actual[i] {
			return false
		}
	}
	return true
}
