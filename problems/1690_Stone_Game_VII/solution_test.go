package _690_Stone_Game_VII

import "testing"

func TestStoneGameVII(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := stoneGameVII([]int{5, 3, 1, 4, 2})
		want := 6
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := stoneGameVII([]int{7, 90, 5, 1, 100, 10, 10, 2})
		want := 122
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
