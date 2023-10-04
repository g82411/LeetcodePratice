package _745_Palindrome_Partitioning_IV

import "testing"

func TestCheckPartitioning(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := checkPartitioning("abcbdd")
		want := true
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := checkPartitioning("bcbddxy")
		want := false
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
