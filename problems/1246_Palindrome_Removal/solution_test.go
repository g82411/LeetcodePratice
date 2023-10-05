package _246_Palindrome_Removal

import "testing"

func TestMinimumMoves(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := minimumMoves([]int{1, 2})
		want := 2
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := minimumMoves([]int{1, 3, 4, 1, 5})
		want := 3
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
