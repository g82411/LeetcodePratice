package _866_Beautiful_Towers_II

import "testing"

func TestMaximumSumOfHeights(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := maximumSumOfHeights([]int{5, 3, 4, 1, 1})
		want := int64(13)
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
