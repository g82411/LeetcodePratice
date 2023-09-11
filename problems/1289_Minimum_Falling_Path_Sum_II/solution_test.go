package _289_Minimum_Falling_Path_Sum_II

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		got := minFallingPathSum(grid)
		want := 13
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
