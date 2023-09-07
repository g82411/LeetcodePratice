package _31_Minimum_Falling_Path_Sum

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	// test case
	t.Run("Case 1", func(t *testing.T) {
		A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		got := minFallingPathSum(A)
		want := 12
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		A := [][]int{{-19, 57}, {-40, -5}}
		got := minFallingPathSum(A)
		want := -59
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})

}
