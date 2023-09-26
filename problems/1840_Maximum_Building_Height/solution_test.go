package _840_Maximum_Building_Height

import "testing"

func TestMaxBuilding(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := maxBuilding(5, [][]int{{2, 1}, {4, 1}})
		want := 2
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := maxBuilding(6, [][]int{})
		want := 5
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
