package _425_Constrained_Subsequence_Sum

import "testing"

func TestConstrainedSubsetSum(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := constrainedSubsetSum([]int{10, 2, -10, 5, 20}, 2)
		want := 37
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := constrainedSubsetSum([]int{-1, -2, -3}, 1)
		want := -1
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
