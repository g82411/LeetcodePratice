package _46_Remove_Boxes

import "testing"

func TestRemoveBox(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1})
		want := 23
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := removeBoxes([]int{1, 1, 1})
		want := 9
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
