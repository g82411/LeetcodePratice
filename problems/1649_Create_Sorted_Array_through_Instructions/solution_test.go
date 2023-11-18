package _649_Create_Sorted_Array_through_Instructions

import "testing"

func TestCreateSortedArray(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := createSortedArray([]int{1, 5, 6, 2})
		want := 1
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := createSortedArray([]int{1, 2, 3, 6, 5, 4})
		want := 3
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
