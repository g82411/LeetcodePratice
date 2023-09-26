package _16_Remove_Duplicate_Letters

import "testing"

func TestRemoveDuplicateLetters(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := removeDuplicateLetters("bcabc")
		want := "abc"
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := removeDuplicateLetters("cbacdcbc")
		want := "acdb"
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
