package _39_Word_Break

import "testing"

func TestWordBreak(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := true
		actual := wordBreak("leetcode", []string{"leet", "code"})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := true
		actual := wordBreak("applepenapple", []string{"apple", "pen"})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		expect := false
		actual := wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 4", func(t *testing.T) {
		expect := true
		actual := wordBreak("aaaaaaa", []string{"aaaa", "aaa"})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
