package _10_Split_Array_Largest_Sum

import "testing"

func TestSplitArray(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := splitArray([]int{7, 2, 5, 10, 8}, 2)
		want := 18
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := splitArray([]int{1, 2, 3, 4, 5}, 2)
		want := 9
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
