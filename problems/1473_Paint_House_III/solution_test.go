package _473_Paint_House_III

import "testing"

func TestMinCost(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		got := minCost([]int{0, 0, 0, 0, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3)
		want := 9
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case2", func(t *testing.T) {
		got := minCost([]int{0, 2, 1, 2, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3)
		want := 11
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
