package _657_Determine_if_Two_Strings_Are_Close

import "testing"

func TestCloseStrings(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		word1 := "abc"
		word2 := "bca"
		expect := true
		actual := closeStrings(word1, word2)
		if actual != expect {
			t.Errorf("word1: %v, word2: %v, expect: %v, actual: %v", word1, word2, expect, actual)
		}
	})

	t.Run("test2", func(t *testing.T) {
		word1 := "a"
		word2 := "aa"
		expect := false
		actual := closeStrings(word1, word2)
		if actual != expect {
			t.Errorf("word1: %v, word2: %v, expect: %v, actual: %v", word1, word2, expect, actual)
		}
	})
}
