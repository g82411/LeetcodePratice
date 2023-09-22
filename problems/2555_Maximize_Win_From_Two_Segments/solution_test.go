package _555_Maximize_Win_From_Two_Segments

import "testing"

func TestMaximizeWin(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := maximizeWin([]int{1, 1, 2, 2, 3, 3, 5}, 2)
		want := 7
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 1", func(t *testing.T) {
		got := maximizeWin([]int{1, 2, 3, 4}, 0)
		want := 2
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
