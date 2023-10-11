package _4_Wildcard_Matching

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := isMatch("aa", "a")
		want := false
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := isMatch("aa", "*")
		want := true
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
