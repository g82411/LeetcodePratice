package _76_Out_of_Boundary_Paths

import "testing"

func TestFindPaths(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := findPaths(2, 2, 2, 0, 0)
		want := 6
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := findPaths(1, 3, 3, 0, 1)
		want := 12
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
