package _74_Dungeon_Game

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}})
		want := 7
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := calculateMinimumHP([][]int{{0}})
		want := 1
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
