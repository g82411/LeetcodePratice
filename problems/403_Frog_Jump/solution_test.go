package _03_Frog_Jump

import "testing"

func TestCanCross(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		got := canCross([]int{0, 1, 3, 5, 6, 8, 12, 17})
		want := true
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case2", func(t *testing.T) {
		got := canCross([]int{0, 1, 2, 3, 4, 8, 9, 11})
		want := false
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
