package _12_Burst_Balloons

import "testing"

func TestMaxCoins(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := maxCoins([]int{3, 1, 5, 8})
		want := 167
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := maxCoins([]int{1, 5})
		want := 10
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
